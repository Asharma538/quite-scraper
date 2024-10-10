import Head from 'next/head';
import { useState } from 'react'
import styles from '../styles/Home.module.css';

export default function Home() {

  const [users, setusers] = useState([]);

  setInterval(() => {
    
  }, 1000);

  return (
    <div className={styles.container}>
      <Head>
        <title>Create Next App</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main>
        <h1 className={styles.title}>
          Lets <a href="https://github.com/Asharma538/quite-scraper">Focus!</a>
        </h1>

        <p className={styles.description}>
          Add the users below whom you want to monitor
        </p>
        {
          users.forEach((idx,elem)=>{
            return <div className={styles.card}>
                      hi
                   </div>
          })
        }
        <div className={styles.card}>
          hi
        </div>
        
      </main>

      <style jsx global>{`
        html,
        body {
          background-color: black;
          padding: 0;
          margin: 0;
          font-family:
            -apple-system,
            BlinkMacSystemFont,
            Segoe UI,
            Roboto,
            Oxygen,
            Ubuntu,
            Cantarell,
            Fira Sans,
            Droid Sans,
            Helvetica Neue,
            sans-serif;
            color: white;
        }
        * {
          box-sizing: border-box;
        }
      `}</style>
    </div>
  );
}
