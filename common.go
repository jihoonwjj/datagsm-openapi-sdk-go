package datagsm

// Ptr는 값 v에 대한 포인터를 반환하는 제네릭 헬퍼입니다.
// 선택적 필드를 인라인으로 초기화할 때 유용합니다.
//
//	datagsm.Ptr(1)        // *int
//	datagsm.Ptr("hello")  // *string
//	datagsm.Ptr(true)     // *bool
func Ptr[T any](v T) *T {
	return &v
}
