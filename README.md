# Reading-Planner-API
Main API designed specifically for the project **Reading-Planner**. The project's objective is to track the user on the books that are currently being read and the ones that are on the queue.
This API will follow the *REST* design pattern, due to its low complexity overall.


## API Design
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

