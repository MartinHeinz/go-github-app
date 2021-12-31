for ARGUMENT in "$@"
do

    KEY=$(echo $ARGUMENT | cut -f1 -d=)
    VALUE=$(echo $ARGUMENT | cut -f2 -d=)

    case "$KEY" in
            APP_ID)           APP_ID=${VALUE} ;;
            INSTALLATION_ID)  INSTALLATION_ID=${VALUE} ;;
            WEBHOOK_SECRET)   WEBHOOK_SECRET=${VALUE} ;;
            REGISTRY)         REGISTRY=${VALUE} ;;
            *)
    esac
done

DUMMY_REGISTRY='ghcr.io/martinheinz/go-github-app'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "\n${BLUE}Renaming variables and files...${NC}\n"

sed -i s~12345~$APP_ID~g cmd/app/utils/utils.go
sed -i s~123456789~$INSTALLATION_ID~g cmd/app/utils/utils.go
sed -i s~WEBHOOK_SECRET~$WEBHOOK_SECRET~g config/server.yaml
sed -i s~$DUMMY_REGISTRY~$REGISTRY~g Makefile

cat $KEY_PATH > config/github-app.pem

echo -e "\n${BLUE}Installing dependencies...${NC}\n"
go mod vendor

echo -e "\n${BLUE}Testing if everything works...${NC}\n"

echo -e "\n${BLUE}Test: make test${NC}\n"
make test
echo -e "\n${BLUE}Test: make build${NC}\n"
make build

# Usage:
# ./configure_project.sh APP_ID="54321" INSTALLATION_ID="987654321" WEBHOOK_SECRET="verysecret" KEY_PATH="./github_key.pem" REGISTRY="ghcr.io/<GITHUB_USERNAME>/go-github-app"
