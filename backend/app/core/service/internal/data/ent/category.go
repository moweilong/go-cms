// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"go-cms/app/core/service/internal/data/ent/category"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Category is the model entity for the Category schema.
type Category struct {
	config `json:"-"`
	// ID of the ent.
	// id
	ID uint32 `json:"id,omitempty"`
	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`
	// 更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`
	// 删除时间
	DeleteTime *int64 `json:"delete_time,omitempty"`
	// 分类名
	Name *string `json:"name,omitempty"`
	// 链接别名
	Slug *string `json:"slug,omitempty"`
	// 描述
	Description *string `json:"description,omitempty"`
	// 缩略图
	Thumbnail *string `json:"thumbnail,omitempty"`
	// 密码
	Password *string `json:"password,omitempty"`
	// 完整路径
	FullPath *string `json:"full_path,omitempty"`
	// 父分类ID
	ParentID *uint32 `json:"parent_id,omitempty"`
	// 优先级
	Priority *int32 `json:"priority,omitempty"`
	// 博文计数
	PostCount    *uint32 `json:"post_count,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Category) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case category.FieldID, category.FieldCreateTime, category.FieldUpdateTime, category.FieldDeleteTime, category.FieldParentID, category.FieldPriority, category.FieldPostCount:
			values[i] = new(sql.NullInt64)
		case category.FieldName, category.FieldSlug, category.FieldDescription, category.FieldThumbnail, category.FieldPassword, category.FieldFullPath:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Category fields.
func (c *Category) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case category.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = uint32(value.Int64)
		case category.FieldCreateTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				c.CreateTime = new(int64)
				*c.CreateTime = value.Int64
			}
		case category.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				c.UpdateTime = new(int64)
				*c.UpdateTime = value.Int64
			}
		case category.FieldDeleteTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_time", values[i])
			} else if value.Valid {
				c.DeleteTime = new(int64)
				*c.DeleteTime = value.Int64
			}
		case category.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = new(string)
				*c.Name = value.String
			}
		case category.FieldSlug:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field slug", values[i])
			} else if value.Valid {
				c.Slug = new(string)
				*c.Slug = value.String
			}
		case category.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				c.Description = new(string)
				*c.Description = value.String
			}
		case category.FieldThumbnail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field thumbnail", values[i])
			} else if value.Valid {
				c.Thumbnail = new(string)
				*c.Thumbnail = value.String
			}
		case category.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				c.Password = new(string)
				*c.Password = value.String
			}
		case category.FieldFullPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field full_path", values[i])
			} else if value.Valid {
				c.FullPath = new(string)
				*c.FullPath = value.String
			}
		case category.FieldParentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field parent_id", values[i])
			} else if value.Valid {
				c.ParentID = new(uint32)
				*c.ParentID = uint32(value.Int64)
			}
		case category.FieldPriority:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field priority", values[i])
			} else if value.Valid {
				c.Priority = new(int32)
				*c.Priority = int32(value.Int64)
			}
		case category.FieldPostCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field post_count", values[i])
			} else if value.Valid {
				c.PostCount = new(uint32)
				*c.PostCount = uint32(value.Int64)
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Category.
// This includes values selected through modifiers, order, etc.
func (c *Category) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// Update returns a builder for updating this Category.
// Note that you need to call Category.Unwrap() before calling this method if this Category
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Category) Update() *CategoryUpdateOne {
	return NewCategoryClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Category entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Category) Unwrap() *Category {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Category is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Category) String() string {
	var builder strings.Builder
	builder.WriteString("Category(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	if v := c.CreateTime; v != nil {
		builder.WriteString("create_time=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := c.UpdateTime; v != nil {
		builder.WriteString("update_time=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := c.DeleteTime; v != nil {
		builder.WriteString("delete_time=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := c.Name; v != nil {
		builder.WriteString("name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := c.Slug; v != nil {
		builder.WriteString("slug=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := c.Description; v != nil {
		builder.WriteString("description=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := c.Thumbnail; v != nil {
		builder.WriteString("thumbnail=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := c.Password; v != nil {
		builder.WriteString("password=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := c.FullPath; v != nil {
		builder.WriteString("full_path=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := c.ParentID; v != nil {
		builder.WriteString("parent_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := c.Priority; v != nil {
		builder.WriteString("priority=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := c.PostCount; v != nil {
		builder.WriteString("post_count=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Categories is a parsable slice of Category.
type Categories []*Category
