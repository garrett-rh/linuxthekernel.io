import { Link, Outlet } from "react-router-dom";

function App() {
  return (
    <>
      <div className="container-fluid">
        <div className="row">
          <nav className="navbar navbar-expand-sm bg-body-tertiary">
            <Link to="/">
              <button className="btn">Home</button>
            </Link>
            <button
              className="navbar-toggler"
              type="button"
              data-bs-toggle="collapse"
              data-bs-target="#navbarNav"
              aria-controls="navbarNav"
              aria-expanded="false"
              aria-label="Toggle navigation"
            >
              <span className="navbar-toggler-icon"></span>
            </button>
            <div className="collapse navbar-collapse" id="navbarNav">
              <ul className="navbar-nav">
                <Link to="/about">
                  <li className="nav-item">
                    <button className="btn">About</button>
                  </li>
                </Link>
                <Link to="/blog">
                  <li className="nav-item">
                    <button className="btn">Blog</button>
                  </li>
                </Link>
              </ul>
            </div>
          </nav>
          <hr />
        </div>
      </div>
      <div className="container">
        <Outlet />
      </div>
    </>
  );
}

export default App;
