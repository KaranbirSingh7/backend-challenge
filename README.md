## Backend Engineer - Coding challenge

This test is a part of our hiring process at crew.work for the Backend Engineer position. It should take you between 2 and 4 hours depending on your experience.

**Feel free to apply! Drop us a line with your LinkedIn/GitHub/Twitter and link of your code repository at [mohamed@crew.work](mailto:mohamed@crew.work)**

At Crew, we use mainly Golang but you are free to use any modern language/framework you wish. Here are some technologies we are more familiar with: JavaScript, Python, Java, Ruby.

> Threat this project as if you/we would continue working on this after your assignment: maintainability, scalability, and readability are super important.

## Goal

For this assignment, you are supposed:
- Load the list of talents exposed on the below API to a MongoDB database.
- Build a RESTful API based on microservices architecture. The API consists of the following endpoints:
  - Endpoint to list all talents from database.
  - Endpoint to add a new talent to the MongoDB database. 
- Write a CI/CD pipeline (CircleCI, GitHub Actions, etc) to test, build and deloy the API.

When you’re done, host it somewhere (e.g. on AWS Lambda, Amazon EC2/EKS, Heroku, Google AppEngine, etc.)

### API documentation

To fetch the list of talents, you’ll need to make request to a publicly-available API to get JSON content:

```json
GET https://hiring.crew.work/v1/talents
```

The above endpoints will return hundreds of talents. You may break down or paginate the results into chunks to make the response easier to handle.

```json
GET https://hiring.crew.work/v1/talents?page=2&limit=20
```

The JSON response contains an array of candidates, each candidate has the following JSON attributes:

```json
{
        "id": "brr63adh5s5m7veepna0",
        "firstName": "Mohamed",
        "lastName": "Labouardy",
        "picture": "https://domain.com/picture.png",
        "job": "CTO @crew.work",
        "location": "Paris, France",
        "linkedin": "https://www.linkedin.com/in/mlabouardy",
        "github": "https://github.com/@mlabouardy",
        "twitter": "https://medium.com/@mlabouardy",
        "tags": [
            "Go",
            "Typescript",
            "AWS"
        ],
        "stage": "🤝 Onsite Interview"
}
```

### Submission

Create a new repo into your favorite git platform (GitHub, Gitlab, etc) with a README file containing list of instructions to run the project.

After you've finished, you can share the repository URL with us.

### Review

After you delivered the completed assignment to us, we will review it as soon as we can, generally within 48 hours. **We pay special attention to:**

- [ ] Coding skills.
- [ ] Usage of AWS, Docker or Serverless framework.
- [ ] Code organization (modularity, dependencies between modules, naming, etc)
- [ ] Overall code quality (edge cases, usage of tools, performance, best practices)

### Good luck,
