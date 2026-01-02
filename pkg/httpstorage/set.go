package httpstorage

//
//import (
//	"github.com/gofiber/fiber/v2"
//)
//
//
//func Set(c *fiber.Ctx, key string) HTTPStorage {
//	data := c.Locals(key)
//	return HTTPStorage{
//		data: data,
//		err:  nil,
//	}
//}
//
//func (h HTTPStorage) String() (string, error) {
//	d := h.data
//	strD, ok := d.(string)
//	if !ok {
//		return "", h.err
//	}
//	return strD, nil
//}
//
//func (h HTTPStorage) Number() (int, error) {
//	d := h.data
//	strD, ok := d.(int)
//	if !ok {
//		return -1, h.err
//	}
//	return strD, nil
//}
