E:\code\Go\work3\ToDoList
├─go.mod
├─go.sum
├─main.go
├─treer.md
├─service
|    ├─UserService
|    |      ├─login.go
|    |      ├─model.go
|    |      └register.go
|    ├─TaskService
|    |      ├─check.go
|    |      ├─create.go
|    |      ├─delete.go
|    |      ├─list.go
|    |      ├─model.go
|    |      ├─search.go
|    |      └update.go
|    ├─AdminService
|    |      ├─addAdmin.go
|    |      ├─blockUser.go
|    |      ├─listUsers.go
|    |      └model.go
├─routes
|   └routes.go
├─pkg
|  ├─e
|  | └e.go
|  ├─authorization
|  |       └jwt.go
├─model
|   ├─model.go
|   ├─password.go
|   ├─task.go
|   └user.go
├─handler
|    ├─userHandler
|    |      ├─idLogin.go
|    |      ├─login.go
|    |      └register.go
|    ├─taskHandler
|    |      ├─check.go
|    |      ├─create.go
|    |      ├─delete.go
|    |      ├─list.go
|    |      ├─search.go
|    |      └update.go
|    ├─adminHandler
|    |      ├─addAdmin.go
|    |      ├─blockUser.go
|    |      ├─listUsers.go
|    |      └login.go
├─docs
|  ├─docs.go
|  ├─swagger.json
|  └swagger.yaml
├─Dao
|  ├─databaseIni.go
|  └migrate.go
├─conf
|  ├─config.go
|  └config.ini