import Groups from "@/components/groups/Groups";
import GroupForm from "@/components/groups/GroupForm";
import styles from "./page.module.css"

function GroupsPage () {
  return (
    <div className={styles.groupPage}>
      <GroupForm />
      <Groups />
    </div>
  )
}

export default GroupsPage;