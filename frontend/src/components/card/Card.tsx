import "./Cards.css";

const Card = () => {
  return (
    <div className="auth-card">
      <h1>Login</h1>
      <div className="input-group">
        <label htmlFor="email">Email</label>
        <input type="email" id="email" placeholder="Enter your email" />
      </div>
      <div className="input-group">
        <label htmlFor="password">Password</label>
        <input
          type="password"
          id="password"
          placeholder="Enter your password"
        />
      </div>
      <button className="login-button" type="submit">
        Login
      </button>
    </div>
  );
};

export default Card;
