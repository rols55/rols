"use client";
import { usePathname } from 'next/navigation'
import GroupEvent from "@/components/events/Event"

function Page () {
  const pathname  = usePathname()

  return (
    <>
      <GroupEvent path={pathname}/>
    </>
  )
}

export default Page;