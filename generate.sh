protoc_version=`protoc --version`

# Enforce protobuf version, this version should be same as the vendored github.com/golang/protobuf package
go_protobuf_version=`git --git-dir="$GOPATH/src/github.com/golang/protobuf/.git" rev-parse HEAD`
# 1. grep vendor.json for the golang/protobuf/proto package, -A1 flag selects the line after the match
# 2. grep only the line that looks like -> "revision": "abcd",
# 3. sed to match -> abcd
# 4. xargs to remove whitespace
# go_protobuf_version_expected=$(cat ${root_dir}vendor/vendor.json | grep -A1 'github.com/golang/protobuf/proto' | grep revision | sed 's/"revision": "\(.*\)",$/\1/' | xargs)
go_protobuf_version_expected="b4deda0973fb4c70b50d226b1af49f3da59f5265"
if [[ $go_protobuf_version != $go_protobuf_version_expected ]]; then
    echo "Version mismatch github.com/golang/protobuf: got ${go_protobuf_version}, expected ${go_protobuf_version_expected}"
    echo "Temporarily switching expected version of github.com/golang/protobuf: ${go_protobuf_version_expected}"
    git --git-dir=$GOPATH/src/github.com/golang/protobuf/.git fetch
    git --git-dir=$GOPATH/src/github.com/golang/protobuf/.git --work-tree=$GOPATH/src/github.com/golang/protobuf/ checkout ${go_protobuf_version_expected}
    (cd $GOPATH/src/github.com/golang/protobuf && make install)
fi

protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative    proto/login.proto