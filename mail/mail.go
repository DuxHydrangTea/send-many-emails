package mail

import (
	"os"
	"strconv"
	"github.com/go-mail/mail/v2"
	"log"
	"sync"
)

type EmailJob struct {
	To string
	Subject string
	Body string
}

func SendEmail(to, subject, body string) error {
	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	m := mail.NewMessage()
	m.SetHeader("From", os.Getenv("SMTP_USER"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	dialer := mail.NewDialer(
		os.Getenv("SMTP_HOST"),
		port,
		os.Getenv("SMTP_USER"),
		os.Getenv("SMTP_PASS"),
	)

	return dialer.DialAndSend(m)
}

func worker(id int, jobs <-chan EmailJob, wg *sync.WaitGroup){
    defer wg.Done()

    port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))

    dialer := mail.NewDialer(
        os.Getenv("SMTP_HOST"),
        port,
        os.Getenv("SMTP_USER"),
        os.Getenv("SMTP_PASS"),
    )

    connect := func() mail.SendCloser {
        s, err := dialer.Dial()
        if err != nil {
            log.Printf("[Worker %d] Failed to connect SMTP: %v", id, err)
            return nil
        }
        log.Printf("[Worker %d] Connected to SMTP", id)
        return s
    }

    s := connect()
    if s == nil {
        return 
    }
    defer s.Close()

    for job := range jobs {
        m := mail.NewMessage()
        m.SetHeader("From", os.Getenv("SMTP_USER"))
        m.SetHeader("To", job.To)
        m.SetHeader("Subject", job.Subject)
        m.SetBody("text/html", job.Body)

        maxRetries := 2
        for i := 0; i < maxRetries; i++ {
            err := mail.Send(s, m)
            if err == nil {
                log.Printf("[Worker %d] Sent to %s", id, job.To)
                break 
            }

            log.Printf("[Worker %d] Error sending to %s: %v. Retrying...", id, job.To, err)
            
            s.Close()
            
            s = connect()
            if s == nil {
                log.Printf("[Worker %d] Failed to reconnect, skipping job for %s", id, job.To)
                break
        }
    }
}


func SendMassEmail(users []string){
	const workCount = 20

	jobs := make(chan EmailJob, len(users))

	var wg sync.WaitGroup

	for w := 1; w <= workCount; w++{
		wg.Add(1)
		go worker(w, jobs, &wg)
	}

	for _, email := range users {
		jobs <- EmailJob{
			To: email,
			Subject: "Ey cuuuuuuuu",
			Body: "Carm on ban da dang kys",
		}
	}

	close(jobs)
	wg.Wait()
}