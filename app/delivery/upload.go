package delivery

import (
	"cloud.google.com/go/storage"
	rsp "github.com/postmy/app/pkg/response"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
	"io"
	"io/ioutil"
	"log"
	"net/url"
)

// Download handler
func (impl *Deliveries) Download() {
	ctx := impl.Ctx
	bucket := "pos-predict.appspot.com"

	ctxt := appengine.NewContext(ctx.Request)
	storageClient, err := storage.NewClient(ctxt, option.WithCredentialsFile("/app/conf/keys.json"))
	if err != nil {
		log.Fatal(err)
		return
	}

	sw, err := storageClient.Bucket(bucket).Object("Proposed Additional Field.png").NewReader(ctxt)
	if err := sw.Close(); err != nil {
		log.Println(err)
		impl.Ctx.WriteString(err.Error())
		return
	}
	x, _ := ioutil.ReadAll(sw)
	impl.Ctx.Output.Body(x)
}

//chmodFile
//func chmodFile() {
//	err := os.Chmod("../conf/keys.json", 0755)
//	if err != nil {
//		log.Fatal(err)
//	}
//}

// Upload handler
func (impl *Deliveries) Upload() {
	ctx := impl.Ctx
	// ctx.WriteString("aa")
	//chmodFile()
	bucket := "pos-predict.appspot.com"

	file, header, err := impl.Controller.GetFile("file")
	if err != nil {
		log.Fatal(err)
		return
	}

	ctxt := appengine.NewContext(ctx.Request)
	storageClient, err := storage.NewClient(ctxt, option.WithCredentialsFile("/app/conf/keys.json"))
	if err != nil {
		log.Fatal(err)
		return
	}

	sw := storageClient.Bucket(bucket).Object(header.Filename).NewWriter(ctxt)
	if _, err := io.Copy(sw, file); err != nil {
		log.Fatal(err)
		return
	}

	if err := sw.Close(); err != nil {
		log.Println(err)
		impl.Ctx.WriteString(err.Error())
		return
	}

	u, err := url.Parse("/" + sw.Attrs().Name)
	if err != nil {
		impl.Ctx.WriteString(err.Error())
		return
	}

	// impl.Ctx.Output.Body(u.EscapedPath())
	//ctx.WriteString(u.EscapedPath())
	rsp.WriteResponse(&impl.Controller, nil, u.EscapedPath())
}
