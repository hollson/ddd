// Copyright 2022 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package services

type TavernOption func(os *Tavern) error

func WithOrderService(os *OrderService) TavernOption {
	return func(t *Tavern) error {
		t.OrderService = os
		return nil
	}
}
