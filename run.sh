search="...YOUR...CREDENTIALS...JSON...CONTENT..."
replace="$(cat backend/creds.json)"
sed  "s~${search}~${replace}~g" docker-compose.yml | docker-compose -f - up