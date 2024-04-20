import React, { useState, useEffect } from 'react'
import { useSearchParams } from 'next/navigation'
import styles from "./page.module.css";

export default function SectionSelector({ tabs }) {
    const [isLoading, setIsLoading] = useState(true);
    const [selectedTab, setSelectedTab] = useState(0);
    const searchParams = useSearchParams()
    
    useEffect(() => {
      const referrer = searchParams.get('referrer')
      if (referrer) {
        for (let i = 0; i < tabs.length; i++) {
          if (tabs[i]?.ref === referrer) {
            setSelectedTab(i);
            break;
          }
        }
      }
      setIsLoading(false);
    }, [searchParams]);

   
    const handleButtonClick = (tab, idx) => {
      setSelectedTab(idx);
        if (tab.callback) {
          tab.callback();
        }
        
        // make this more NextJS?
        if (history.pushState) {
          var newurl = window.location.protocol + "//" + window.location.host + window.location.pathname + '?referrer='+tabs[idx].ref;
          window.history.pushState({path:newurl},'',newurl);
      }
    };

    if (isLoading) {
      return <div>Loading...</div>;
    }

    return (
      <>
        <div className={styles.buttonRow}>
        { tabs ? tabs.map( (tab, idx) => (
          tab ? 
            <button key={idx}
            onClick={() => handleButtonClick(tab, idx)}
            className={styles.button27 + " " +(idx === selectedTab ? styles.selected : null)}
            >
                {tab.name}
            </button>
          : null
        )) : null}
        </div>
        <div className={styles.tabContent}>
            {tabs[selectedTab]?.content ? tabs[selectedTab].content() : null}
        </div>
      </>
    )
}
