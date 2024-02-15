rm -rf mocks/mock_*.go

bin/mockgen -source globalping/client.go -destination mocks/mock_client.go -package mocks
bin/mockgen -source view/viewer.go -destination mocks/mock_viewer.go -package mocks
bin/mockgen -source utils/time.go -destination mocks/mock_time.go -package mocks
