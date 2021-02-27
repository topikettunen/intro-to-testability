# PolarSquad - Intro to Testing 

Software projects tend to label testing under a large umbrella of different testing definitions. We often see tests like unit tests, acceptance tests, integration tests, system tests, and so many more. These can get quite confusing relatively quickly, especially to those who don't work primarily in the field of quality assurance.

I'm by no means a testing expert, even though I've worked in a QA company for a few years, where I could see how these things are done. I worked as a DevOps consultant there, but you can often see these two being linked together quite closely.

## Principle of testing

The goal of testing is really to make the software or system to fail. By achieving this, we know that there are defects in the system. Testing reduces the number of defects, but at the same time, it is impossible to test all different functionalities in the systems exhaustively. However, unfortunately, business people want to see 100% test coverage always. Good luck with that.

The big idea behind testing is finding defects and bugs as early as possible since they are the cheapest to fix. It also tends to be the case that defects seem to cluster up in the system. Most of the system bugs and defects are found in a small set of modules or functionalities.

Testing is also this weird practice in itself that doesn't trust successes in your system. It's a fallacy to believe that your system might be absent of error, so usually, by running your tests repeatedly, if you're fixing these errors, eventually, all the tests will pass. This doesn't mean that mistakes are gone. If you happen to fix any of the code, you might often add new code to it, which might cause some issues. This is why tests need constant reviewing to ensure that they keep on testing the right things in the future.

## Defining levels of testing

Before we get started on how we might approach this topic from the practical perspective, I think we could set some basic definitions between these many different testing levels.

### Micro-tests and unit tests

Unit tests can be seen in many software projects as this habit or practice is something that programmers have to do. Quite often, people tend to miss why these can be very beneficial for the whole project. People also often confuse these with some other form of testing and hence miss unit tests themselves.

In its most straightforward form unit test is a test that ensures that the production code does what the programmer expects it to do. People tend to see unit tests as tests written for a little scope, e.g. function or a group of functions. The latter one is often called "micro-tests", a term coined by Mike Hill.

An important thing to note in these tests is that these tests should not be coupled with any other aspect of the system/program other than that small part that the micro or unit tests might be testing. People often tend to call unit testing an awful idea or practice due to their fragile nature. This fragile nature often comes when these tests are unnecessarily coupled with parts of the system that they shouldn't, e.g. databases, servers, etc. This will cause these tests to be fragile. When we start to link these parts together, we get into the integration testing territory, which we discuss in the following detail.

These tests are also used mainly for rapid feedback to the programmer, so they shouldn't take a long time to run. Usually, if this part of testing takes a long time, the tests themselves are poorly written, or the tests do something that they shouldn't do. We discuss how to make your program to be testable later.

### Functional tests

Functional tests can be seen as a unit test but written at a more extensive scope. Appropriate mocks are often used for different slow parts of the system if needed (e.g. mock APIs, mock databases).

### Integration tests
    
Micro tests and unit tests are often considered the foundation for all the testing. Every other level of testing is built upon that. Micro tests and unit tests are also always the tests that the programmer should write. After we can ensure that the smallest building blocks of the program work, we can integrate them into other systems in which integration tests come in.

These tests are often done by architects or, for example, technical leads to ensure that the sub-assembly of system components works correctly together. An important thing to notice here is that these are just "plumbing" tests, as in connecting things. Business rules and expectations don't yet apply to these tests.

### System tests (end-to-end)

System tests are then the next level in terms of integrations. These are also integration tests, in a sense, as they ensure that the internal parts of the system are integrated and configured correctly. Again these tests are "plumbing" tests, and business rules don't apply to these.

### Acceptance tests

Acceptance tests are tests written by the business to ensure that the written code works the way it wants it to work. These tests are not written by technical people anymore. Or the technical people who might write these represent the business, for example, QA engineer or business analyst.

## Defining methods and approaches to testing

Usually, testing methods can be split into static and dynamic testing. The difference between the two is that static testing doesn't execute the code where dynamic testing does.

On the other hand, testing approaches can be split into white box and black box testing, and sometimes "gray box" testing. White box testing is a way of testing the code to know how the code is written. This kind of testing is quite often done at a unit level, for example, by the programmer. On the other hand, black-box testing removes the code's knowledge from testing, which results in just testing the system's behaviour or, for example, which its input-output system works. 

## Keeping tests clean

In software projects, tests are often strongly coupled with production code, which is a big pain point when we talk about these tests' fragility. To get your tests to be not so strongly coupled is decoupling them might take some time and can be relatively difficult depending on the system. You often don't want to use too many mocks in your tests, and at the same, you don't want to couple those tests to something unnecessary. This is why writing useful tests requires a little bit of design skills, especially object-oriented, and usage of different design patterns.

Tests often also tend to become pretty fragile since they are not treated the same way as the rest of the system. You can see various ad hoc styles of doing these tests, which happens to be convenient for the writer. Then programmers also tend to think that mocking is somewhat of a vital part of testing. It sometimes is, but also not always. There seems to be a trend amongst programmers that almost everything should be mocked, at least by the looks of how many different mock libraries there are.

I don't know where this idea comes from that mocking should be almost always done. Still, I believe that programmers tend to focus on the wrong parts of the testing, for example, focusing only on a small scope of micro tests where some little bit more functional approach would be more straightforward, and often better since many of the aspects that micro tests focus on tend to be tested indirectly with other tests. 

### Writing testable code

An important skill when writing code is the ability to make your code testable. Quite often, programmers don't necessarily think of this aspect when they happen in the programming zone. To write a testable code is also quite often a merit of good code, since if you happen to utilize different aspects of what some could call bad code in your codebase, that makes the testing much harder.

#### SOLID Principles

SOLID principles are various methods of trying to make software design more maintainable and readable. Robert C. Martin first introduced these principles in his 2000 paper Design Principles and Design Patterns. SOLID  is an acronym for the following principles:

The Single-responsibility principle: a class should only have a single responsibility; that is, only changes to one part of the software's specification should be able to affect the specification of the class.
The Open–closed principle: "software entities ... should be open for extension, but closed for modification."
The Liskov substitution principle: "objects in a program should be replaceable with instances of their subtypes without altering the correctness of that program." See also design by contract.
The Interface segregation principle: "many client-specific interfaces are better than one general-purpose interface."
The Dependency inversion principle: "depend upon abstractions, not concretions."

These principles were mainly designed for object-oriented programming, but similar ideas can be utilized in general software development. The following example is difficult to test, but we can fix it quite nicely with these principles.

``` go
var db *sql.DB

...

func writeUserName(id int) {
	var name string
	err := db.Query("select name from users where id = ?", id).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	f, err = os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err = f.WriteString(name)
	if err != nil {
		log.Fatal(err)
	}
}
```

Here you can see a code example that is tightly coupled to a database and the file system. What happens when the database doesn't exist or isn't running? Or when `data.txt` is already open or just cannot be modified? To make it a little bit more testable, we need to be able to break these couplings.

``` go
type myDB struct {
	db *sql.DB
}

func newMyDB(db *sql.DB) *UserRepo {
	return &myDB{
		db: db,
	}
}

func (db *myDB) writeUserName(id int, file string) {
	var name string
	err := db.Query("select name from users where id = ?", id).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	f, err = os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err = f.WriteString(name)
	if err != nil {
		log.Fatal(err)
	}
}
```

This is one way to make it a little bit more testable. This also allows us to some better mocking against possible DB queries. Now there are multiple ways of doing this, so I'm just mainly making a point here.

#### Non-determinism in code

Very simple example of non-testable code is different non-deterministic factors in your code. Let's look at this example:

``` go
func getTimeOfDay() {
    now := time.Now()
    if now.Hour >= 0 && now.Hour < 6 {
	    return "Night"
    }
    if now.Hour >= 6 && now.Hour < 12 {
	    return "Morning"
    }
    if now.Hour >= 12 && now.Hour < 18 {
	    return "Afternoon"
    }
    return "Evening"
}
```

Considering the functionality of this function, it works as it should. It gets the current time of day. Unfortunately, in this form, the code is not testable unless we want to start tinkering with the system's time itself. This is also a great example of tightly coupling in the software. It is not possible to reuse this method for possible date and time processing in any way. It has these hidden inputs inside, so it lies about the information required to get the day's time. Thankfully fixing this is easy:

``` go
func getTimeOfDay(now Time) {
    if now.Hour >= 0 && now.Hour < 6 {
	    return "Night"
    }
    if now.Hour >= 6 && now.Hour < 12 {
	    return "Morning"
    }
    if now.Hour >= 12 && now.Hour < 18 {
	    return "Afternoon"
    }
    return "Evening"
}
```

Time is an excellent example of non-determinism. Even though we fixed this on this one level, we have just moved the problem to another abstraction level. So we need to care about the same thing on that level too. This makes time also a great example of where mockups can be very useful. I'm not going into detail about how this is done since the point is to picture how this kind of non-determinism might affect your code's testability.

## Conclusion

So there is a lot to cover when we talk about testing. Fortunately, software project quite often tends to have dedicated people for doing most of this stuff. Still, when it comes to this practice's practicality, it is the programmers who are doing it. Especially object-oriented way of thinking causes much cognitive overhead when you need to start thinking testability of the software. This may seem like a lot of work, but almost always, it is necessary for the overall success of the project.

Although... in functional programming, this kind of testability often comes free when you embrace that paradigm. Many of these tips on testability are mainly related to the mutability of this or that. So functional thinking is a good perk if you want to be able to create good testable code. Even though you might be working with popular languages like Java or C++, good software design and craftsmanship are worthy goals for any program.
