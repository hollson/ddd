// Copyright (c) 2014 - The Event Horizon authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package guestlist

import (
	eh "github.com/looplab/eventhorizon"
)

// 事件包含「类型」和「数据」
const (
	InviteCreatedEvent   eh.EventType = "InviteCreated"   // 创建邀请事件
	InviteAcceptedEvent  eh.EventType = "InviteAccepted"  // 邀请接受事件
	InviteDeclinedEvent  eh.EventType = "InviteDeclined"  // 邀请谢绝事件
	InviteConfirmedEvent eh.EventType = "InviteConfirmed" // 邀请确认事件
	InviteDeniedEvent    eh.EventType = "InviteDenied"    // 请求拒绝事件
)

// 注册事件数据
func init() {
	eh.RegisterEventData(InviteCreatedEvent, func() eh.EventData { return &InviteCreatedData{} })
}

// InviteCreatedData 事件数据
type InviteCreatedData struct {
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}
