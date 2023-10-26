Hello, this is what i can do with the time provided for this test

✅ database is created and persistance (mysql & auto migration from gorm) <br />
✅ user can login using phone number / email (with hashed password) <br />
✅ product is retrieved from database <br />
✅ checkout and stock deduction is running <br />
❕ i will separate between business lock and table lock, from what i understand, this case only need business lock if queue system can be used, as table lock (oversold case) will be resolved if the stock deduction is queued <br />
✅ business lock (delete if no payment is made) is running <br />
❕ i use goroutine as the release stock mechanism, but this is a bad practice as it will spawn many goroutines every request, i will recommend using scheduled task if the time is enough <br /><br />

❓ i cant get the corelation between the shop and orders, because if orders go to the shop then product should connect to warehouse not the order<br /><br />

❌ no test case is generated<br />
❌ warehouse service is not started yet<br />
❌ no integration for all of the microservices as i didn't learn kubernetes from scratch yet<br />
❕ horizontal scaling would be possible if this project used kubernetes<br />
❌ no robust error handling yet, only log.Print is used<br />
❌ no monitoring and logging yet
