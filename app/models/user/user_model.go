// package models
// @author: CongYuKun
// @since: 2023/1/5
// @desc: user_model

package user

import "gohub/app/models"

type User struct {
    models.BaseModel

    Name     string `json:"name,omitempty"`
    Email    string `json:"-"`
    Phone    string `json:"-"`
    Password string `json:"-"`

    models.CommonTimestampsField
}
