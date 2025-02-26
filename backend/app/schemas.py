from pydantic import BaseModel

class Token(BaseModel):
    access_token: str
    token_type: str

class TokenData(BaseModel):
    username: str | None = None

class User(BaseModel):
    username: str
    email: str | None = None
    hashed_password: str

class UserInDB(User):
    username: str
    email: str | None = None
    hashed_password: str