import os
import argparse
from dotenv import load_dotenv
from pathlib import Path
from google import genai
from google.genai import types
from prompts import system_prompt

def main():
    print("Hello from ai-agent-python!")
    load_dotenv()
    api_key = os.getenv("GEMINI_API_KEY")
    print(api_key)

    if api_key is None:
        raise RuntimeError("GEMINI_API_KEY is not set in the environment variables.")

    parser = argparse.ArgumentParser(description="A simple AI agent using Gemini API.")
    parser.add_argument("user_prompt", type=str, help="User Prompt for the AI agent.", default="Why are episodes 7-9 so much worse than 1-6? Use one paragraph.")
    parser.add_argument("--verbose", action="store_true", help="Enable verbose output")
    args = parser.parse_args()
    # access `args.user_prompt`` 
    
    messages: list[types.Content] = [
      types.Content(role="user", parts=[types.Part(text=args.user_prompt)])
    ]
    
    config=types.GenerateContentConfig(
        system_instruction=system_prompt,
        temperature=0
    )

    client = genai.Client(api_key=api_key)
    response = client.models.generate_content(
        model="gemini-2.5-flash",
        contents=messages,
        config=config,
    )
    
    if args.verbose:
      print(f"User prompt: {messages[0].parts[0].text}")
      print(f"Prompt tokens: {response.usage_metadata.prompt_token_count}")
      print(f"Response tokens: {response.usage_metadata.candidates_token_count}")
      print(f"Response: {response.text}")
    else:
      print(response.text)

if __name__ == "__main__":
    main()
