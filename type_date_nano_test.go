package carbon_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dromara/carbon/v2"
)

func TestDateNano_Scan(t *testing.T) {
	c := carbon.NewDateNano(carbon.Now())

	t.Run("[]byte type", func(t *testing.T) {
		assert.Nil(t, c.Scan([]byte(carbon.Now().ToDateNanoString())))
	})

	t.Run("string type", func(t *testing.T) {
		assert.Nil(t, c.Scan(carbon.Now().ToDateNanoString()))
	})

	t.Run("int64 type", func(t *testing.T) {
		assert.Nil(t, c.Scan(carbon.Now().Timestamp()))
	})

	t.Run("time type", func(t *testing.T) {
		assert.Nil(t, c.Scan(carbon.Now().StdTime()))
	})

	t.Run("unsupported type", func(t *testing.T) {
		assert.Error(t, c.Scan(nil))
	})
}

func TestDateNano_Value(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		c := carbon.NewCarbon()
		v, err := carbon.NewDateNano(c).Value()
		assert.Nil(t, v)
		assert.Nil(t, err)
	})

	t.Run("invalid time", func(t *testing.T) {
		c := carbon.Parse("xxx")
		v, err := carbon.NewDateNano(c).Value()
		assert.Nil(t, v)
		assert.Error(t, err)
	})

	t.Run("valid time", func(t *testing.T) {
		c := carbon.Parse("2020-08-05")
		v, err := carbon.NewDateNano(c).Value()
		assert.Equal(t, c.StdTime(), v)
		assert.Nil(t, err)
	})
}

func TestDateNano_MarshalJSON(t *testing.T) {
	type User struct {
		DateNano carbon.DateNano `json:"date_nano"`
	}

	t.Run("zero time", func(t *testing.T) {
		var user User

		c := carbon.NewCarbon()
		user.DateNano = carbon.NewDateNano(c)

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"date_nano":""}`, string(data))
	})

	t.Run("invalid time", func(t *testing.T) {
		var user User

		c := carbon.Parse("xxx")
		user.DateNano = carbon.NewDateNano(c)

		data, err := json.Marshal(&user)
		assert.Error(t, err)
		assert.Empty(t, string(data))
	})

	t.Run("valid time", func(t *testing.T) {
		var user User

		c := carbon.Parse("2020-08-05 13:14:15.999999999")
		user.DateNano = carbon.NewDateNano(c)

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"date_nano":"2020-08-05.999999999"}`, string(data))
	})
}

func TestDateNano_UnmarshalJSON(t *testing.T) {
	type User struct {
		DateNano carbon.DateNano `json:"date_nano"`
	}

	t.Run("empty value", func(t *testing.T) {
		var user User

		value := `{"date_nano":""}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))
		assert.Empty(t, user.DateNano.String())
	})

	t.Run("null value", func(t *testing.T) {
		var user User

		value := `{"date_nano":"null"}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))
		assert.Empty(t, user.DateNano.String())
	})

	t.Run("zero value", func(t *testing.T) {
		var user User

		value := `{"date_nano":"0"}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))
		assert.Empty(t, user.DateNano.String())
	})

	t.Run("valid value", func(t *testing.T) {
		var user User

		value := `{"date_nano":"2020-08-05.999999999"}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))
		assert.Equal(t, "2020-08-05.999999999", user.DateNano.String())
	})
}

func TestDateNano_String(t *testing.T) {
	c := carbon.Parse("2020-08-05 13:14:15.999999999")
	assert.Equal(t, "2020-08-05.999999999", carbon.NewDateNano(c).String())
}

func TestDateNano_GormDataType(t *testing.T) {
	assert.Equal(t, "time", carbon.NewDateNano(carbon.Now()).GormDataType())
}
