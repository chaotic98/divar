# Divar CLI

یک پروژه ساده شبیه‌ساز دیوار با رابط خط فرمان (CLI) نوشته شده با زبان Go

## امکانات

### مدیریت کاربران
- ثبت‌نام کاربر جدید با نام کاربری منحصر به فرد
- احراز هویت کاربران

### مدیریت آگهی‌ها
- ایجاد آگهی جدید با عنوان و تگ‌های دلخواه
- حذف آگهی (فقط توسط مالک آگهی)
- نمایش لیست آگهی‌های کاربر
- فیلتر آگهی‌ها بر اساس تگ

### سیستم علاقه‌مندی‌ها
- افزودن آگهی به لیست علاقه‌مندی‌ها
- حذف آگهی از لیست علاقه‌مندی‌ها
- نمایش لیست علاقه‌مندی‌ها با امکان فیلتر

## پیش‌نیازها

- نصب شده بودن [Go 1.16+](https://golang.org/dl/)
- ترمینال یا خط فرمان

## نصب و راه‌اندازی

1. کلون کردن ریپازیتوری:

```bash
git clone https://github.com/your-username/divar-cli.git
cd divar-cli
```
2.کامپایل و اجرای پروژه: 
```bash
go build -o divar-cli
./divar-cli
```

## دستورات

```bash
register <username>  # ثبت کاربر جدید
add_advertise <username> <title> [tag]  # افزودن آگهی جدید
rem_advertise <username> <title>       # حذف آگهی
list_my_advertises <username> [tag]    # نمایش آگهی‌های من
add_favorite <username> <title>             # افزودن به علاقه‌مندی‌ها
rem_favorite <username> <title>             # حذف از علاقه‌مندی‌ها
list_favorite_advertises <username> [tag]   # نمایش علاقه‌مندی‌ها
```

## مثال‌های عملی

در این بخش چند مثال کاربردی از نحوه کار با Divar CLI را مشاهده می‌کنید:

### ثبت کاربر و افزودن آگهی
```bash
register user1
register user2
add_advertise user1 "خودروی صفر کیلومتر" "وسایل نقلیه"
add_advertise user2 "لپ‌تاپ گیمینگ" "کامپیوتر"
add_advertise user2 "موبایل نو" "دیجیتال"
```
