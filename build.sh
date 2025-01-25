echo GO_BLOG
cd ./go-blog
GOOS=linux GOARCH=amd64 go build -o bin/backend .


echo GO_ECOM
cd ../go-ecom
GOOS=linux GOARCH=amd64 go build -o bin/backend .


echo GO_GATEWAY
cd ../go-gateway
GOOS=linux GOARCH=amd64 go build -o bin/backend .


echo GO_SASS
cd ../go-sass
GOOS=linux GOARCH=amd64 go build -o bin/backend .


cd ..

docker-compose up
# ls

# git add .
# git commit -m "DEPLOY:BUILD"