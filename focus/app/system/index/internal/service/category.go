package service

import (
	"context"
	"fmt"
	"focus/app/dao"
	"focus/app/model"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/os/gcache"
	"github.com/gogf/gf/util/gconv"
	"time"
)

// 栏目管理服务
var Category = categoryService{}

type categoryService struct{}

const (
	mapCacheKey       = "category_map_cache"
	mapCacheDuration  = time.Hour
	treeCacheKey      = "category_tree_cache"
	treeCacheDuration = time.Hour
)

var (
	// 缩进字符串，用于层级展示
	indentChars = []string{"&nbsp;", " │", " ├", " └"}
)

// 查询列表
func (s *categoryService) GetTree(ctx context.Context, contentType string) ([]*model.CategoryTreeItem, error) {
	v, err := gcache.GetOrSetFunc(treeCacheKey+contentType, func() (interface{}, error) {
		entities, err := s.GetList(ctx)
		if err != nil {
			return nil, err
		}
		tree, err := s.formTree(0, contentType, entities)
		if err != nil {
			return nil, err
		}
		return tree, nil
	}, treeCacheDuration)
	if err != nil {
		return nil, err
	}
	return v.([]*model.CategoryTreeItem), nil
}

// 获取指定栏目ID及其下面所有子ID，构成数组返回。
// 注意，返回的ID列表中包含查询的栏目ID.
func (s *categoryService) GetSubIdList(ctx context.Context, id uint) ([]uint, error) {
	m, err := s.GetMap(ctx)
	if err != nil {
		return nil, err
	}
	entity := m[id]
	if entity == nil {
		return nil, gerror.Newf(`%d栏目不存在`, id)
	}
	tree, err := s.GetTree(ctx, entity.ContentType)
	if err != nil {
		return nil, err
	}
	return append([]uint{id}, s.getSubIdListByTree(id, tree)...), nil
}

// 递归获取指定栏目ID下的所有子级
func (s *categoryService) getSubIdListByTree(id uint, trees []*model.CategoryTreeItem) []uint {
	idArray := make([]uint, 0)
	for _, item := range trees {
		if item.ParentId == id {
			idArray = append(idArray, item.Id)
			if len(item.Items) > 0 {
				idArray = append(idArray, s.getSubIdListByTree(item.Id, item.Items)...)
			}
		} else if len(item.Items) > 0 {
			idArray = append(idArray, s.getSubIdListByTree(id, item.Items)...)
		}
	}
	return idArray
}

// 构造树形栏目列表。
func (s *categoryService) formTree(parentId uint, contentType string, entities []*model.Category) ([]*model.CategoryTreeItem, error) {
	tree := make([]*model.CategoryTreeItem, 0)
	for _, entity := range entities {
		if contentType != "" && entity.ContentType != contentType {
			continue
		}
		if entity.ParentId == parentId {
			subTree, err := s.formTree(entity.Id, contentType, entities)
			if err != nil {
				return nil, err
			}
			item := &model.CategoryTreeItem{
				Items: subTree,
			}
			if err = gconv.Struct(entity, item); err != nil {
				return nil, err
			}
			tree = append(tree, item)
		}
	}
	return tree, nil
}

// 获得所有的栏目列表。
func (s *categoryService) GetList(ctx context.Context) ([]*model.Category, error) {
	orderBy := fmt.Sprintf(
		`%s ASC, %s ASC`,
		dao.Category.Columns.Sort,
		dao.Category.Columns.Id,
	)
	return dao.Category.Order(orderBy).All()
}

// 查询单个栏目信息
func (s *categoryService) GetItem(ctx context.Context, id uint) (*model.Category, error) {
	m, err := s.GetMap(ctx)
	if err != nil {
		return nil, err
	}
	return m[id], nil
}

// 获得所有的栏目列表，构成Map返回，键名为栏目ID
func (s *categoryService) GetMap(ctx context.Context) (map[uint]*model.Category, error) {
	v, err := gcache.GetOrSetFunc(mapCacheKey, func() (interface{}, error) {
		entities, err := s.GetList(ctx)
		if err != nil {
			return nil, err
		}
		m := make(map[uint]*model.Category)
		for _, entity := range entities {
			item := entity
			m[entity.Id] = item
		}
		return m, nil
	}, mapCacheDuration)
	if err != nil {
		return nil, err
	}
	return v.(map[uint]*model.Category), nil
}
