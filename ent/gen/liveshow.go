// Code generated by ent, DO NOT EDIT.

package gen

import (
	"backend/ent/gen/liveshow"
	"backend/ent/gen/upload"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// LiveShow is the model entity for the LiveShow schema.
type LiveShow struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// Legend holds the value of the "legend" field.
	Legend string `json:"legend,omitempty"`
	// Duration holds the value of the "duration" field.
	Duration int `json:"duration,omitempty"`
	// Viewed holds the value of the "viewed" field.
	Viewed bool `json:"viewed,omitempty"`
	// CurrentPlayed holds the value of the "currentPlayed" field.
	CurrentPlayed bool `json:"currentPlayed,omitempty"`
	// StartedTime holds the value of the "started_time" field.
	StartedTime time.Time `json:"started_time,omitempty"`
	// EndedTime holds the value of the "ended_time" field.
	EndedTime time.Time `json:"ended_time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LiveShowQuery when eager-loading is set.
	Edges            LiveShowEdges `json:"edges"`
	live_show_upload *int
	selectValues     sql.SelectValues
}

// LiveShowEdges holds the relations/edges for other nodes in the graph.
type LiveShowEdges struct {
	// Upload holds the value of the upload edge.
	Upload *Upload `json:"upload,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UploadOrErr returns the Upload value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LiveShowEdges) UploadOrErr() (*Upload, error) {
	if e.Upload != nil {
		return e.Upload, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: upload.Label}
	}
	return nil, &NotLoadedError{edge: "upload"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*LiveShow) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case liveshow.FieldViewed, liveshow.FieldCurrentPlayed:
			values[i] = new(sql.NullBool)
		case liveshow.FieldID, liveshow.FieldDuration:
			values[i] = new(sql.NullInt64)
		case liveshow.FieldLegend:
			values[i] = new(sql.NullString)
		case liveshow.FieldCreateTime, liveshow.FieldStartedTime, liveshow.FieldEndedTime:
			values[i] = new(sql.NullTime)
		case liveshow.ForeignKeys[0]: // live_show_upload
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the LiveShow fields.
func (ls *LiveShow) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case liveshow.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ls.ID = int(value.Int64)
		case liveshow.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				ls.CreateTime = value.Time
			}
		case liveshow.FieldLegend:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field legend", values[i])
			} else if value.Valid {
				ls.Legend = value.String
			}
		case liveshow.FieldDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration", values[i])
			} else if value.Valid {
				ls.Duration = int(value.Int64)
			}
		case liveshow.FieldViewed:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field viewed", values[i])
			} else if value.Valid {
				ls.Viewed = value.Bool
			}
		case liveshow.FieldCurrentPlayed:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field currentPlayed", values[i])
			} else if value.Valid {
				ls.CurrentPlayed = value.Bool
			}
		case liveshow.FieldStartedTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field started_time", values[i])
			} else if value.Valid {
				ls.StartedTime = value.Time
			}
		case liveshow.FieldEndedTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field ended_time", values[i])
			} else if value.Valid {
				ls.EndedTime = value.Time
			}
		case liveshow.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field live_show_upload", value)
			} else if value.Valid {
				ls.live_show_upload = new(int)
				*ls.live_show_upload = int(value.Int64)
			}
		default:
			ls.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the LiveShow.
// This includes values selected through modifiers, order, etc.
func (ls *LiveShow) Value(name string) (ent.Value, error) {
	return ls.selectValues.Get(name)
}

// QueryUpload queries the "upload" edge of the LiveShow entity.
func (ls *LiveShow) QueryUpload() *UploadQuery {
	return NewLiveShowClient(ls.config).QueryUpload(ls)
}

// Update returns a builder for updating this LiveShow.
// Note that you need to call LiveShow.Unwrap() before calling this method if this LiveShow
// was returned from a transaction, and the transaction was committed or rolled back.
func (ls *LiveShow) Update() *LiveShowUpdateOne {
	return NewLiveShowClient(ls.config).UpdateOne(ls)
}

// Unwrap unwraps the LiveShow entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ls *LiveShow) Unwrap() *LiveShow {
	_tx, ok := ls.config.driver.(*txDriver)
	if !ok {
		panic("gen: LiveShow is not a transactional entity")
	}
	ls.config.driver = _tx.drv
	return ls
}

// String implements the fmt.Stringer.
func (ls *LiveShow) String() string {
	var builder strings.Builder
	builder.WriteString("LiveShow(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ls.ID))
	builder.WriteString("create_time=")
	builder.WriteString(ls.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("legend=")
	builder.WriteString(ls.Legend)
	builder.WriteString(", ")
	builder.WriteString("duration=")
	builder.WriteString(fmt.Sprintf("%v", ls.Duration))
	builder.WriteString(", ")
	builder.WriteString("viewed=")
	builder.WriteString(fmt.Sprintf("%v", ls.Viewed))
	builder.WriteString(", ")
	builder.WriteString("currentPlayed=")
	builder.WriteString(fmt.Sprintf("%v", ls.CurrentPlayed))
	builder.WriteString(", ")
	builder.WriteString("started_time=")
	builder.WriteString(ls.StartedTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("ended_time=")
	builder.WriteString(ls.EndedTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// LiveShows is a parsable slice of LiveShow.
type LiveShows []*LiveShow
