import apiFetch from "../apiFetch";
export default function Edit({ profile, setIsOpen }) {
  async function handleSubmit(data) {
    //take the data and apply it to profile object on the same keys
    for (let [key, value] of data.entries()) {
      profile[key] = value;
    }
    profile.public = profile.public === "on" ? true : false;
    await apiFetch("/profile/edit", () => setIsOpen(false), {
      method: "POST",
      body: JSON.stringify(profile),
    });
  }
  return (
    <>
      <div className="profile-dialog">
        <form action={handleSubmit} className="form">
          <div>
            <img
              className="profile-avatar"
              src={`http://${process.env.NEXT_PUBLIC_API_HOST_CLIENT_SIDE}:${process.env.NEXT_PUBLIC_API_PORT_CLIENT_SIDE}/api/getimageuser/${profile.image}`}
              alt="User avatar"
            />
          </div>
          <label>Gender</label>
          <div>
            <select name="sex" defaultValue={profile.sex}>
              <option>male</option>
              <option>female</option>
            </select>
          </div>

          <label>Name</label>
          <div>
            <input
              type="text"
              name="firstname"
              defaultValue={profile.firstname}
            />
          </div>

          <label>Surname</label>
          <div>
            <input
              type="text"
              name="lastname"
              defaultValue={profile.lastname}
            />
          </div>

          <label>Nickname</label>
          <div>
            <input
              type="text"
              name="nickname"
              defaultValue={profile.nickname}
              className="form-control"
            />
          </div>

          <label>E-mail address</label>
          <div>
            <input
              type="email"
              name="email"
              className="form-control"
              defaultValue={profile.email}
            />
          </div>

          <label>About me</label>
          <div>
            <textarea
              rows="3"
              name="aboutme"
              className="form-control form-about"
              defaultValue={profile.aboutme}
            ></textarea>
          </div>

          <div className="checkbox">
            <input
              type="checkbox"
              name="public"
              defaultChecked={profile.public}
            />
            <label htmlFor="checkbox_1">Make this account public</label>
          </div>

          <div className="button-container">
            <button type="action" className="button-27">
              Save
            </button>
            <button
              type="button"
              className="button-27"
              onClick={() => setIsOpen(false)}
            >
              Back
            </button>
          </div>
        </form>
      </div>
    </>
  );
}
