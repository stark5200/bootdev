import os
import argparse
import sys
from dotenv import load_dotenv
from prompts import system_prompt
from call_function import available_functions, call_function
from openai import OpenAI

def main():
    print("Hello from ai-agent-python!")
    load_dotenv()
    api_key = os.getenv("OPENROUTER_API_KEY")

    if api_key is None:
        raise RuntimeError("OPENROUTER_API_KEY is not set in the environment variables.")

    parser = argparse.ArgumentParser(description="A simple AI agent using OpenRouter API.")
    parser.add_argument("user_prompt", type=str, help="User Prompt for the AI agent.")
    parser.add_argument("--verbose", action="store_true", help="Enable verbose output")
    args = parser.parse_args()
    client = OpenAI(
        base_url="https://openrouter.ai/api/v1",
        api_key=api_key,
    )
    
    messages = [
        {"role": "system", "content": system_prompt},
        {"role": "user", "content": args.user_prompt},
    ]
    
    for _ in range(20):
        result = run_agent(messages, client, args)
        if result:
            print("Final response:")
            print(result)
            break
    else:
        print("Max iterations reached without a final response.")
        sys.exit(1)
        
def run_agent(messages: list, client: OpenAI, args: argparse.Namespace) -> list:    
    response = client.chat.completions.create(
        model="openrouter/free",
        messages=messages,
        tools=available_functions,
        temperature=0,
    )
    
    message = response.choices[0].message
    
    if args.verbose:
        print(f"User prompt: {args.user_prompt}")
        print(f"Prompt tokens: {response.usage.prompt_tokens}")
        print(f"Response tokens: {response.usage.completion_tokens}")
    
    if message.tool_calls:
        for tool_call in message.tool_calls:
            result_message = call_function(tool_call, args.verbose)
            if not result_message["content"]:
                raise Exception("Function call returned no content.")
            if args.verbose:
                print(f"-> {result_message['content']}")
    else:
        print(message.content)
        print("No tool calls detected. Agent finished.")
        return messages
        
    messages.append(message)
    messages.append(result_message)
    return None


if __name__ == "__main__":
    main()
