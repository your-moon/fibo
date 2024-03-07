<h1 align="center">
    <img height="80" width="160" src="./assets/gopher-icon.gif" alt="Go"><br>Skill sharing platform

</h1>

> Skill sharing platform

### Үндсэн функционал шаардлага

- [x] Системд бүртгүүлсний дараа мэдээ submit хийх боломжтой
- [x] Мэдээ бичих үед code snippet оруулах, Italic, Bold гэх мэт уншууртай
      байлгах компонент, tool-үүд, энгийн WYSIWYG - постлогдсоны дараа ч
      яг харагдаж байгаа шигээ байдлаар edit хийх ёстой.
- [x] Backoffice дээр тухай post-г уншиж review хийсний дараа Approve хийх
      үед үндсэн сайт дээр пост-логдоно
- [x] Мэдээн дээр Бүртгүүлсэн хэрэглэгч болон Anynomous хэрэглэгч Like
      дарах болон коммент бичих боломжтой байна. Энэ нь тухайн author-н
      reputation point-д нөлөөлнө.
- [ ] Reputation point нь тухайн author мэдээ тус бүрээс хэдийн цалин
      авахыг шийдвэрлэх ба reputation point өндөр байх тусам өндөр цалин
      авна.
- [ ] Цалинг сарын төгсгөлд бодож author тус бүрт олгох дүнг тооцоолно.
- [ ] Нийт author болон цалин, reputation point зэргийг харж, удирдах хэсэг
      backoffice хэсэгт байна.
- [x] Мэдээ удирдлагын хэсэг мөн backoffice хэсэгт байна.
- [x] Author нь өөрийн dashboard дээрээс мөн ойр зуурын статистик
      мэдээллийг харж, оруулсан мэдээний жагсаалт Нийтлэгдсэн/
      Хүлээгдэж байгаа гэх зэргээр ангилж хардаж байна.
- [x] Мэдээ удирдлагын хэсэг мөн backoffice хэсэгт байна.

### Нэмэлт шаардлага

- [x] Веб нь responsive байх
- [x] Үг үсгийн алдаа, UI/UX нь гажиггүй байх, Цэвэрхэн хялбар загвартай
      байх
- [ ] Нийтлэл тус бүрийг уншсан хүний тоолуур байх

### Техникийн шаардлага

- [x] Frontend ReactJS ашиглана. Бэлэн template, component ашиглаж
      болно.
- [x] Backend, Database технологийг өөрийн хүслээр сонгох. Бэлэн tool-үүд
      болон код ашиглаж болно.

### Нэмэлтээр оруулж болох функцууд

- [ ] Үг тоолж хичнээн минут уншихад харуулах (Medium шиг).
- [x] Мэдээний категори байх - Категорийг хэрхэн зохион байгуулах
      удирдах зэргийг өөрөө мэдэж хийнэ үү.
- [x] Мэдээг устгах, өөрчлөх боломж
- [ ] Мэдээний дор коммент бичсэн үед түүнийг reply хийх боломжтой байх
- [ ] Post-г social сувгууд дээр шэйр хийж болдог байх
- [ ] Буруу мэдээлэл эсвэл ёс бус мэдээ байх үед report хийх боломжтой
      байх
- [ ] Платформд хэрэгтэй өөр бусад боломжуудыг нэмж оруулах
      (Чөлөөтэй сэтгээд хийнэ үү)
- [ ] Үнэгүй SSL тохируулах
- [x] Docker контейнэр болгох
- [x] Frontend дээр SSR ашиглах

### Backend Technologies

- [x] Programming Language: Go
- [x] HTTP Framework: Gin
- [x] Database: PostgreSQL
- [x] Database ORM: goqu
- [x] Validation & Parsing: govalidator & kong

### Frontend Technologies

- [x] Programming Language: Typescript
- [x] Web Framework: NextJS
- [x] Front-end Library: NextUI
- [x] WYSIWYG Editor: EditorJS
- [x] Editor Parser: editorjs-parser & editorjs-html

## Makefile

Makefile requires installed dependecies:

- [go](https://go.dev/doc/install)
- [docker-compose](https://docs.docker.com/compose/reference)
- [migrate](https://github.com/golang-migrate/migrate)

```shell
$ make

Usage: make [command]

Commands:
 rename-project name={name}    Rename project

 build-http                    Build http server

 migration-create name={name}  Create migration
 migration-up                  Up migrations
 migration-down                Down last migration

 docker-up                     Up docker services
 docker-down                   Down docker services

 fmt                           Format source code
 test                          Run unit tests

```

## HTTP Server

```shell
$ ./bin/http-server --help

Usage: http-server

Flags:
  -h, --help               Show context-sensitive help.
      --env-path=STRING    Path to env config file
```

**Configuration** is based on the environment variables. See [.env.template](./config/env/.env.template).

```shell
# Expose env vars before and start server
$ ./bin/http-server

# Expose env vars from the file and start server
$ ./bin/http-server --env-path ./config/env/.env
```

## License

This project is licensed under the [MIT License](https://github.com/pvarentsov/fibo/blob/main/LICENSE).
