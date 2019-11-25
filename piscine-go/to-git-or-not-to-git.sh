
curl -s https://01.alem.school/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{githubLogin:{_eq:\"Abaisky\"}}){id}}"}' | jq '.data.user[].id'