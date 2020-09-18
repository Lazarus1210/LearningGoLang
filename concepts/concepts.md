What are contexts in golang

//An empty context is a Context that has no value, no deadline and it’s never canceled. The context.Background() function returns a default empty Context. This Context is generally used to derive other context objects since it never cancels. It can also be used in test cases or merely to pass a context object to an API where custom context is not important.


        ctx, cancel := context.WithCancel(context.Background()) <------------ this is how you create it 

        go square(ctx) // start square goroutine

