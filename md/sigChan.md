Explanation of graceful shutdown code from cmd/main.go

```
sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	sig := <-sigChan
	product_logger.Println("RECEIVED TERMINATE, Graceful shutdown", sig)

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(30)*time.Second)
	server.Shutdown(ctx)
```

Signal Notification: The program creates a channel named sigChan to receive OS signals. This is done using `signal.Notify(sigChan, os.Interrupt, os.Kill)`, which configures the program to relay incoming interrupt `(Ctrl+C)` and kill signals to sigChan. This setup allows the program to react to these signals for a controlled shutdown.

Waiting for a Signal: The statement `sig := <-sigChan` is where the channel's power is truly showcased. This line blocks the execution of the program until a signal is received on sigChan. It's a clear example of how channels can synchronize the flow of a program — in this case, pausing the program's execution until it's time to shut down.

Graceful Shutdown: Once a signal is received, the program proceeds to shut down the server gracefully. It uses `context.WithTimeout` to create a context with a deadline, which is then passed to `server.Shutdown(ctx)`. This ensures that the server has a chance to finish processing any ongoing requests before it shuts down, rather than terminating abruptly.
