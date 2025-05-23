class Bot:
    def __init__(self):
        # This list will hold all the registered command handlers
        self.command_handlers = []

    def message_handler(self, commands=None):
        # This is the actual decorator
        def decorator(func):
            # Store a handler with its filter (commands) and callback function
            self.command_handlers.append({
                "commands": commands,
                "callback": func
            })
            return func  # Required so the decorated function still works normally
        return decorator

    def receive_message(self, message_text):
        # Simulate scanning incoming messages and matching them to a handler
        if not message_text.startswith("/"):
            print("Not a command.")
            return

        command = message_text[1:].split()[0]  # Extract command, e.g. "/start" → "start"
        for handler in self.command_handlers:
            if handler["commands"] and command in handler["commands"]:
                print(f"Dispatching to handler for command: /{command}")
                handler["callback"](message_text)
                return

        print(f"No handler found for: /{command}")

# Instantiate the bot
bot = Bot()

# Register handlers using decorators
@bot.message_handler(commands=['start', 'hello'])
def send_welcome(message):
    print(f"Welcome handler triggered! Message: {message}")

@bot.message_handler(commands=['help'])
def send_help(message):
    print("Help handler triggered!")

# Simulate incoming messages
bot.receive_message("/start")
bot.receive_message("/help")
bot.receive_message("/unknown")
