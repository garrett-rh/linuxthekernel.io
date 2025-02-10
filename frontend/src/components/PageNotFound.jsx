import { Link } from "react-router-dom";

function PageNotFound() {
    return (
        <div className="container">
            <h1 className="d-flex justify-content-center">Page Not Found</h1>
            <Link to="/" className="d-flex justify-content-center"><button>Go Home</button></Link>
        </div>
    );
}

export default PageNotFound;
