package levenshtein

import "testing"

func Test_levenshtein(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "horse->ros",
			args: args{"horse", "ros"},
			want: 3,
		},
		{
			name: "intention->execution",
			args: args{"intention", "execution"},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levenshtein(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("levenshtein() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCompute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Compute(`{
  "_id": {
    "$oid": "624a4d2913c23571d47ec445"
  },
  "id": 2,
  "name": "sadasdas",
  "description": "asdasdas",
  "owner": {
    "id": 51322445,
    "login": "peanutzhen",
    "avatar_url": "https://avatars.githubusercontent.com/u/51322445?v=4",
    "email": "astzls213@gmail.com"
  },
  "is_running": false,
  "traffic_config": {
    "device": "lo0",
    "port": 8080,
    "addr": "localhost:8081",
    "onlineaddr": ""
  },
  "filter_config": {
    "http_path_regex_filter": null
  },
  "advance_config": {
    "is_recursion_diff": false,
    "is_ignore_array_sequence": false
  },
  "total_record": 3,
  "success_record": 3,
  "created_time": "2022-04-04 09:43:05",
  "updated_time": "2022-04-04 09:43:05"
}`, `{
  "_id": {
    "$oid": "624c410fd14713fbd5960f2e"
  },
  "id": 4,
  "name": "测试user接口",
  "description": "测试user接口是否符合预期",
  "owner": {
    "id": 51322445,
    "login": "peanutzhen",
    "avatar_url": "https://avatars.githubusercontent.com/u/51322445?v=4",
    "email": "astzls213@gmail.com"
  },
  "is_running": true,
  "traffic_config": {
    "device": "lo0",
    "port": 8080,
    "addr": "localhost:8081",
    "onlineaddr": "127.0.0.1:8080"
  },
  "filter_config": {
    "http_path_regex_filter": null
  },
  "advance_config": {
    "is_recursion_diff": false,
    "is_ignore_array_sequence": true
  },
  "total_record": 7,
  "success_record": 6,
  "created_time": "2022-04-05 21:15:59",
  "updated_time": "2022-04-06 15:11:18"
}`)
	}
}
