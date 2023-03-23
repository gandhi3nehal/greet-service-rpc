protoc ./api/v1/*.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --proto_path=.

https://github.com/gandhi3nehal/greet-service-rpc.git

…or create a new repository on the command line
echo "# greet-service-rpc" >> README.md
git init
git add README.md
git commit -m "first commit"
git branch -M main
git remote add origin https://github.com/gandhi3nehal/greet-service-rpc.git
git push -u origin main
…or push an existing repository from the command line
git remote add origin https://github.com/gandhi3nehal/greet-service-rpc.git
git branch -M main
git push -u origin main

$ git remote set-url origin https://git-repo/new-repository.git
