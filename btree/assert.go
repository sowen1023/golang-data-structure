/* Copyright © INFINI LTD. All rights reserved.
 * Web: https://infinilabs.com
 * Email: hello#infini.ltd */

package btree

func assert(b bool) {
	if !b {
		panic("invalid parameter")
	}
}
