package httpstorage

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Get(c *fiber.Ctx, key string) HTTPStorage {
	data := c.Locals(key)

	return HTTPStorage{
		data: data,
		err:  nil,
	}
}

func (h HTTPStorage) String() (string, error) {
	d := h.data
	strD, ok := d.(string)
	if !ok {
		return "", fmt.Errorf("an error on read locals")
	}
	return strD, nil
}

func (h HTTPStorage) Number() (int, error) {
	d := h.data
	strD, ok := d.(string)
	if !ok {
		return -1, fmt.Errorf("an error on read locals")
	}
	fi, err := strconv.Atoi(strD)
	if err != nil {
		return -1, err
	}
	return fi, nil
}
