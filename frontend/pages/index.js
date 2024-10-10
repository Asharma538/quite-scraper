import Head from 'next/head';
import { useEffect, useState } from 'react'
import styles from '../styles/Home.module.css';

export default function Home() {

  const [users, setUsers] = useState([]);

  useEffect(() => {
    const interval = setInterval(() => {
      fetch('http://localhost:8080/getactivity')
      .then(response => response.json())
      .then(data => {
        if (data==null) return;
        
        console.log(data);
        setUsers(data["users_with_updates"]);
      })
      .catch(err => {
        console.log(err);
      })
    }, 10000);

    return () => clearInterval(interval);
  }, []);

  return (
    <div className={styles.container}>
      <Head>
        <title>Quite Scraper</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main>
        <h1 className={styles.title}>
          Lets <a href="https://github.com/Asharma538/quite-scraper">Focus!</a>
        </h1>

        <p className={styles.description}>
          See updates from the users you're monitoring
        </p>
        <div>{
          users.map(elem=>{
            return <div className={styles.card}>
              {elem}
            </div>
          })
        }</div>
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
