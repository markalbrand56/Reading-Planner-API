# Reading-Planner-API
![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![MongoDB](https://img.shields.io/badge/MongoDB-4EA94B?style=for-the-badge&logo=mongodb&logoColor=white)

REST API designed for a *reading planner*. The project's objective is to track the user on the books that are currently being read and the ones that are on the queue. It will allow the user to add new books, and track when the user reads them.


## Model Design
For this project, an API must be design to handle all the information about the books. The following will be the basic information handled by the API:

```
{
  ID
  Title
  Author
  Volume
  Progress
}
```

## Database
All the information needs to be stored somehow, so the database used for this project will be:
**Mongo DB**

This database was chosen thanks to its flexibility, ease to use, and the offer of a free plan.
