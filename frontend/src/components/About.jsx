import { Link } from "react-router-dom";
import Fishing from "../../imgs/fishing.svg";

function About() {
  return (
    <>
      <p className="fs-1 text-center">About Me</p>
      <hr />
      <div className="row">
        <div className="col-md-6">
          <img
            src={Fishing}
            className="rounded mx-auto d-block img-thumbnail img-fluid"
            alt="Fishing"
          ></img>
        </div>
        <div className="col-md-6">
          <p>
            I'm Garrett Harvey. I am a DevOps & Software Engineer. This is a
            website that I made for fun. I intend to make this mostly about my
            technical work and personal projects but more than likely it'll just
            end up being blog posts about my time spent fishing.
          </p>
          <Link to="/blog/AboutMe">For more information on who I am and how I started, check this post.</Link>
          <p>You can also find my LinkedIn and GitHub linked in the top right! My most up to date resume can be found on my LinkedIn profile.</p>
        </div>
      </div>
    </>
  );
}

export default About;
