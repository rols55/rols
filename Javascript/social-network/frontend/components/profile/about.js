export default function About({ profile }) {
  return (
    <>
      <table className="info">
        <tbody>
          <tr>
            <td>Username:</td>
            <td>{profile.username}</td>
          </tr>
          <tr>
            <td>Email:</td>
            <td>{profile.email}</td>
          </tr>
          <tr>
            <td>Gender:</td>
            <td>{profile.sex}</td>
          </tr>
          <tr>
            <td>Birthday:</td>
            <td>{profile.birthday}</td>
          </tr>
          <tr>
            <td>Account:</td>
            <td>{profile.public ? "Public" : "Private"}</td>
          </tr>
        </tbody>
      </table>
    </>
  );
}
