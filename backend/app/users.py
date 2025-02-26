from sqlalchemy.orm import Session
from schemas import User, UserInDB
from database import Base, engine
from passlib.context import CryptContext
from sqlalchemy import Column, Integer, String

pwd_context = CryptContext(schemes=["bcrypt"], deprecated="auto")

class UserModel(Base):
    __tablename__= "users"
    id = Column(Integer, primary_key=True, index=True)
    username = Column(String, unique=True, index=True)
    email = Column(String, unique=True, index=True)
    hashed_password = Column(String)

def get_password_hash(password: str) -> str:
    return pwd_context.hash(password)

def verify_password(plain_password: str, hashed_password: str) -> bool:
    return pwd_context.verify(plain_password, hashed_password)

def authenticate_user(db: Session, username: str, password: str) -> UserInDB | None:
    user = get_by_username(db, username)
    if not user or not verify_password(password, user.hashed_password):
        return None
    return user

def get_by_username(db: Session, username: str) -> UserInDB | None:
    return db.query(UserModel).filter(UserModel.username == username).first()