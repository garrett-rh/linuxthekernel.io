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
        </div>
      </div>

      <p> linkedin </p>
      <p> github </p>
    </>
  );
}

export default About;
