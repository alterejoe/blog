# This is what I want to do the next time I work on this project:

## for file upload and processing
need to finish file upload by batching over files on upload submit

need to work on terraform/postgres rls policies 

also need to work on casbin maybe /{county}/{election} 

### RBAC 
Groups have internal roles
Roles do not have internal groups

- within my current schema
*A user must be assigned a group and that group must have a role to access information*
Users are assigned to groups - client(allen), department(sales), company(harp)
Endpoints are assigned to roles - /api -> admin, /proofing -> proofing, /shipping -> shipping
Roles are assigned to groups - admin -> client, proofing -> department, shipping -> company




## things to get done today
- [ ] import contest files into database
- [ ] generate ballot based on terms 
- [ ] create iterative checklist of contests
- [ ] create sqlc for setting app.group_id for rls 
        - this will be added to the start of every transaction or pool connection
        - probably best to put directly into the dep injection functions

## 
- [ ] upload pdfs
- [ ] create iterative checklist of candidates

