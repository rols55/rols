"use client";
import Group from "@/components/groups/Group";
import { useState, useEffect } from 'react';
import { usePathname } from 'next/navigation'
import apiFetch from "@/components/apiFetch";

function Page () {
  const pathname  = usePathname()
  const [resData, setResData] = useState(null)
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        apiFetch(pathname, setResData);
      } catch (error) {
        setError(error);
      } finally {
        setIsLoading(false);
      }
    };

    fetchData();
  }, []);

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error: {error.message}</div>;
  }

  return (
    <>
      <Group res={resData}/>
    </>
  )
}

export default Page;