// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"context"
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
	"time"

	"focus/app/model"
)

// ReplyDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type ReplyDao struct {
	gmvc.M
	DB      gdb.DB
	Table   string
	Columns replyColumns
}

// ReplyColumns defines and stores column names for table gf_reply.
type replyColumns struct {
	Id         string // 回复ID
	ParentId   string // 回复对应的上一级回复ID(没有的话默认为0)
	Title      string // 回复标题
	Content    string // 回复内容
	TargetType string // 评论类型: content, reply
	TargetId   string // 对应内容ID，可能回复的是另一个回复，所以这里没有使用content_id
	UserId     string // 网站用户ID
	ZanCount   string // 赞
	CaiCount   string // 踩
	CreatedAt  string // 创建时间
	UpdatedAt  string //
}

var (
	// Reply is globally public accessible object for table gf_reply operations.
	Reply = ReplyDao{
		M:     g.DB("default").Model("gf_reply").Safe(),
		DB:    g.DB("default"),
		Table: "gf_reply",
		Columns: replyColumns{
			Id:         "id",
			ParentId:   "parent_id",
			Title:      "title",
			Content:    "content",
			TargetType: "target_type",
			TargetId:   "target_id",
			UserId:     "user_id",
			ZanCount:   "zan_count",
			CaiCount:   "cai_count",
			CreatedAt:  "created_at",
			UpdatedAt:  "updated_at",
		},
	}
)

// Ctx is a chaining function, which creates and returns a new DB that is a shallow copy
// of current DB object and with given context in it.
// Note that this returned DB object can be used only once, so do not assign it to
// a global or package variable for long using.
func (d *ReplyDao) Ctx(ctx context.Context) *ReplyDao {
	return &ReplyDao{M: d.M.Ctx(ctx)}
}

// As sets an alias name for current table.
func (d *ReplyDao) As(as string) *ReplyDao {
	return &ReplyDao{M: d.M.As(as)}
}

// TX sets the transaction for current operation.
func (d *ReplyDao) TX(tx *gdb.TX) *ReplyDao {
	return &ReplyDao{M: d.M.TX(tx)}
}

// Master marks the following operation on master node.
func (d *ReplyDao) Master() *ReplyDao {
	return &ReplyDao{M: d.M.Master()}
}

// Slave marks the following operation on slave node.
// Note that it makes sense only if there's any slave node configured.
func (d *ReplyDao) Slave() *ReplyDao {
	return &ReplyDao{M: d.M.Slave()}
}

// Args sets custom arguments for model operation.
func (d *ReplyDao) Args(args ...interface{}) *ReplyDao {
	return &ReplyDao{M: d.M.Args(args...)}
}

// LeftJoin does "LEFT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").LeftJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").LeftJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *ReplyDao) LeftJoin(table ...string) *ReplyDao {
	return &ReplyDao{M: d.M.LeftJoin(table...)}
}

// RightJoin does "RIGHT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").RightJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").RightJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *ReplyDao) RightJoin(table ...string) *ReplyDao {
	return &ReplyDao{M: d.M.RightJoin(table...)}
}

// InnerJoin does "INNER JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").InnerJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").InnerJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *ReplyDao) InnerJoin(table ...string) *ReplyDao {
	return &ReplyDao{M: d.M.InnerJoin(table...)}
}

// Fields sets the operation fields of the model, multiple fields joined using char ','.
// The parameter <fieldNamesOrMapStruct> can be type of string/map/*map/struct/*struct.
func (d *ReplyDao) Fields(fieldNamesOrMapStruct ...interface{}) *ReplyDao {
	return &ReplyDao{M: d.M.Fields(fieldNamesOrMapStruct...)}
}

// FieldsEx sets the excluded operation fields of the model, multiple fields joined using char ','.
// The parameter <fieldNamesOrMapStruct> can be type of string/map/*map/struct/*struct.
func (d *ReplyDao) FieldsEx(fieldNamesOrMapStruct ...interface{}) *ReplyDao {
	return &ReplyDao{M: d.M.FieldsEx(fieldNamesOrMapStruct...)}
}

// Option sets the extra operation option for the model.
func (d *ReplyDao) Option(option int) *ReplyDao {
	return &ReplyDao{M: d.M.Option(option)}
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (d *ReplyDao) OmitEmpty() *ReplyDao {
	return &ReplyDao{M: d.M.OmitEmpty()}
}

// Filter marks filtering the fields which does not exist in the fields of the operated table.
func (d *ReplyDao) Filter() *ReplyDao {
	return &ReplyDao{M: d.M.Filter()}
}

// Where sets the condition statement for the model. The parameter <where> can be type of
// string/map/gmap/slice/struct/*struct, etc. Note that, if it's called more than one times,
// multiple conditions will be joined into where statement using "AND".
// Eg:
// Where("uid=10000")
// Where("uid", 10000)
// Where("money>? AND name like ?", 99999, "vip_%")
// Where("uid", 1).Where("name", "john")
// Where("status IN (?)", g.Slice{1,2,3})
// Where("age IN(?,?)", 18, 50)
// Where(User{ Id : 1, UserName : "john"})
func (d *ReplyDao) Where(where interface{}, args ...interface{}) *ReplyDao {
	return &ReplyDao{M: d.M.Where(where, args...)}
}

// WherePri does the same logic as M.Where except that if the parameter <where>
// is a single condition like int/string/float/slice, it treats the condition as the primary
// key value. That is, if primary key is "id" and given <where> parameter as "123", the
// WherePri function treats the condition as "id=123", but M.Where treats the condition
// as string "123".
func (d *ReplyDao) WherePri(where interface{}, args ...interface{}) *ReplyDao {
	return &ReplyDao{M: d.M.WherePri(where, args...)}
}

// And adds "AND" condition to the where statement.
func (d *ReplyDao) And(where interface{}, args ...interface{}) *ReplyDao {
	return &ReplyDao{M: d.M.And(where, args...)}
}

// Or adds "OR" condition to the where statement.
func (d *ReplyDao) Or(where interface{}, args ...interface{}) *ReplyDao {
	return &ReplyDao{M: d.M.Or(where, args...)}
}

// Group sets the "GROUP BY" statement for the model.
func (d *ReplyDao) Group(groupBy string) *ReplyDao {
	return &ReplyDao{M: d.M.Group(groupBy)}
}

// Order sets the "ORDER BY" statement for the model.
func (d *ReplyDao) Order(orderBy ...string) *ReplyDao {
	return &ReplyDao{M: d.M.Order(orderBy...)}
}

// Limit sets the "LIMIT" statement for the model.
// The parameter <limit> can be either one or two number, if passed two number is passed,
// it then sets "LIMIT limit[0],limit[1]" statement for the model, or else it sets "LIMIT limit[0]"
// statement.
func (d *ReplyDao) Limit(limit ...int) *ReplyDao {
	return &ReplyDao{M: d.M.Limit(limit...)}
}

// Offset sets the "OFFSET" statement for the model.
// It only makes sense for some databases like SQLServer, PostgreSQL, etc.
func (d *ReplyDao) Offset(offset int) *ReplyDao {
	return &ReplyDao{M: d.M.Offset(offset)}
}

// Page sets the paging number for the model.
// The parameter <page> is started from 1 for paging.
// Note that, it differs that the Limit function start from 0 for "LIMIT" statement.
func (d *ReplyDao) Page(page, limit int) *ReplyDao {
	return &ReplyDao{M: d.M.Page(page, limit)}
}

// Batch sets the batch operation number for the model.
func (d *ReplyDao) Batch(batch int) *ReplyDao {
	return &ReplyDao{M: d.M.Batch(batch)}
}

// Cache sets the cache feature for the model. It caches the result of the sql, which means
// if there's another same sql request, it just reads and returns the result from cache, it
// but not committed and executed into the database.
//
// If the parameter <duration> < 0, which means it clear the cache with given <name>.
// If the parameter <duration> = 0, which means it never expires.
// If the parameter <duration> > 0, which means it expires after <duration>.
//
// The optional parameter <name> is used to bind a name to the cache, which means you can later
// control the cache like changing the <duration> or clearing the cache with specified <name>.
//
// Note that, the cache feature is disabled if the model is operating on a transaction.
func (d *ReplyDao) Cache(duration time.Duration, name ...string) *ReplyDao {
	return &ReplyDao{M: d.M.Cache(duration, name...)}
}

// Data sets the operation data for the model.
// The parameter <data> can be type of string/map/gmap/slice/struct/*struct, etc.
// Eg:
// Data("uid=10000")
// Data("uid", 10000)
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
func (d *ReplyDao) Data(data ...interface{}) *ReplyDao {
	return &ReplyDao{M: d.M.Data(data...)}
}

// All does "SELECT FROM ..." statement for the model.
// It retrieves the records from table and returns the result as []*model.Reply.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of M.Where function,
// see M.Where.
func (d *ReplyDao) All(where ...interface{}) ([]*model.Reply, error) {
	all, err := d.M.All(where...)
	if err != nil {
		return nil, err
	}
	var entities []*model.Reply
	if err = all.Structs(&entities); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

// One retrieves one record from table and returns the result as *model.Reply.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of M.Where function,
// see M.Where.
func (d *ReplyDao) One(where ...interface{}) (*model.Reply, error) {
	one, err := d.M.One(where...)
	if err != nil {
		return nil, err
	}
	var entity *model.Reply
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindOne retrieves and returns a single Record by M.WherePri and M.One.
// Also see M.WherePri and M.One.
func (d *ReplyDao) FindOne(where ...interface{}) (*model.Reply, error) {
	one, err := d.M.FindOne(where...)
	if err != nil {
		return nil, err
	}
	var entity *model.Reply
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindAll retrieves and returns Result by by M.WherePri and M.All.
// Also see M.WherePri and M.All.
func (d *ReplyDao) FindAll(where ...interface{}) ([]*model.Reply, error) {
	all, err := d.M.FindAll(where...)
	if err != nil {
		return nil, err
	}
	var entities []*model.Reply
	if err = all.Structs(&entities); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

// Struct retrieves one record from table and converts it into given struct.
// The parameter <pointer> should be type of *struct/**struct. If type **struct is given,
// it can create the struct internally during converting.
//
// The optional parameter <where> is the same as the parameter of Model.Where function,
// see Model.Where.
//
// Note that it returns sql.ErrNoRows if there's no record retrieved with the given conditions
// from table and <pointer> is not nil.
//
// Eg:
// user := new(User)
// err  := dao.User.Where("id", 1).Struct(user)
//
// user := (*User)(nil)
// err  := dao.User.Where("id", 1).Struct(&user)
func (d *ReplyDao) Struct(pointer interface{}, where ...interface{}) error {
	return d.M.Struct(pointer, where...)
}

// Structs retrieves records from table and converts them into given struct slice.
// The parameter <pointer> should be type of *[]struct/*[]*struct. It can create and fill the struct
// slice internally during converting.
//
// The optional parameter <where> is the same as the parameter of Model.Where function,
// see Model.Where.
//
// Note that it returns sql.ErrNoRows if there's no record retrieved with the given conditions
// from table and <pointer> is not empty.
//
// Eg:
// users := ([]User)(nil)
// err   := dao.User.Structs(&users)
//
// users := ([]*User)(nil)
// err   := dao.User.Structs(&users)
func (d *ReplyDao) Structs(pointer interface{}, where ...interface{}) error {
	return d.M.Structs(pointer, where...)
}

// Scan automatically calls Struct or Structs function according to the type of parameter <pointer>.
// It calls function Struct if <pointer> is type of *struct/**struct.
// It calls function Structs if <pointer> is type of *[]struct/*[]*struct.
//
// The optional parameter <where> is the same as the parameter of Model.Where function,
// see Model.Where.
//
// Note that it returns sql.ErrNoRows if there's no record retrieved and given pointer is not empty or nil.
//
// Eg:
// user  := new(User)
// err   := dao.User.Where("id", 1).Scan(user)
//
// user  := (*User)(nil)
// err   := dao.User.Where("id", 1).Scan(&user)
//
// users := ([]User)(nil)
// err   := dao.User.Scan(&users)
//
// users := ([]*User)(nil)
// err   := dao.User.Scan(&users)
func (d *ReplyDao) Scan(pointer interface{}, where ...interface{}) error {
	return d.M.Scan(pointer, where...)
}

// Chunk iterates the table with given size and callback function.
func (d *ReplyDao) Chunk(limit int, callback func(entities []*model.Reply, err error) bool) {
	d.M.Chunk(limit, func(result gdb.Result, err error) bool {
		var entities []*model.Reply
		err = result.Structs(&entities)
		if err == sql.ErrNoRows {
			return false
		}
		return callback(entities, err)
	})
}

// LockUpdate sets the lock for update for current operation.
func (d *ReplyDao) LockUpdate() *ReplyDao {
	return &ReplyDao{M: d.M.LockUpdate()}
}

// LockShared sets the lock in share mode for current operation.
func (d *ReplyDao) LockShared() *ReplyDao {
	return &ReplyDao{M: d.M.LockShared()}
}

// Unscoped enables/disables the soft deleting feature.
func (d *ReplyDao) Unscoped() *ReplyDao {
	return &ReplyDao{M: d.M.Unscoped()}
}
