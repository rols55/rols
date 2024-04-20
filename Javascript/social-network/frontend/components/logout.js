import React from "react";
import { redirect } from "next/navigation";
import { cookies } from "next/headers";

const Logout = () => {
  return (
    <form
      action={async () => {
        "use server";
          const url = `http://${process.env.API_HOST_SERVER_SIDE}:${process.env.API_PORT_SERVER_SIDE}/api/logout`;
          await fetch(url, {credentials: 'include'})
          cookies().delete("session_token");
          redirect("/")
      }}
      >
      <button type="submit">Logout</button>
    </form>
  );
};

export default Logout;
