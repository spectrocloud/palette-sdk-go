from pydantic_settings import BaseSettings


class Settings(BaseSettings):
    tool_server_url: str = "http://localhost:8090"
    host: str = "0.0.0.0"
    port: int = 8080
    crewai_llm: str = "anthropic/claude-sonnet-4-20250514"
    journey_llm: str = "anthropic/claude-opus-4-20250514"
    anthropic_api_key: str = ""
    openai_api_key: str = ""
    log_level: str = "info"

    model_config = {"env_prefix": "", "case_sensitive": False}


settings = Settings()
