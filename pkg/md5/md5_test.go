package md5

import "testing"

func TestSum(t *testing.T) {
	// 测试 Sum 方法
	type args struct {
		data []byte // 输入数据
	}
	tests := []struct {
		name string // 用例名称
		args args   // 输入参数
		want string // 期望结果
	}{
		{
			name: "空字符串",
			args: args{data: []byte("")},
			want: "d41d8cd98f00b204e9800998ecf8427e", // 空字符串的MD5
		},
		{
			name: "hello world",
			args: args{data: []byte("hello world")},
			want: "5eb63bbbe01eeed093cb22bb8f5acdc3",
		},
		{
			name: "中文测试",
			args: args{data: []byte("你好，世界")},
			want: "dbefd3ada018615b35588a01e216ae6e",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sum(tt.args.data)
			if got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
