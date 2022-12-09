-Manual Type Declaration
คือประกาศตัวแปรพร้อมระบุชนิดข้อมูล
โครงสร้าง  var<ชื่อตัวแปร><ชนิดข้อมูล>    ex     var name string
                                            name = "jaofourth"

-Type Inference
คือประกาศตัวแปรโดยไม่ระบุชนิดข้อมูล
โครงสร้าง <ชื่อตัวแปร>:=<value>    ex  name:="jaofourth" 

-รับค่าจากคีย์บอร์ดด้วย Scanf
โครงสร้าง  fmt.Scanf(string_format,address_list)
                    *string_format คือ รูปแบบตัวแทนชนิดข้อมูล
                            string use %s
                            interger use %d
                            floating point use %f
                    *address_list คือ ตัวเก็บข้อมูล

-รูปแบบคำสั่งแบบเงื่อนไขเดียว
if เงื่อนไข{//true
    คำสั่งเมื่อเงื่อนไขเป็นจริง;
}

-if...else statement
if เงื่อนไข{
    คำสั่งเมื่อเงื่อนไขเป็นจริง
}else{
    คำสั่งเมื่อเงื่อนไขเป็นเท็จ 
}

-Switch...Case
switch ค่าที่เก็บในตัวแปรควบคุม{
    case ค่าที่1:คำสั่งที่1
    case ค่าที่2:คำสั่งที่2
    case ค่าที่n:คำสั่งที่n
    default:คำสั่งเมื่อไม่มีค่าตรงกับที่ระบุในcase
}

-Array

var number[4] int = [4]int{10,20,30,40}
len(number)

names :=[2]string{"สำรวย","สำเร็จ"}
len(names)

-Slice
names :=[]srting{"สำรวย","สำเร็จ"}
//วิธีเพิ่มสมาชิกลงในslice
names = append(names,"สมภพ")
names = append(names,"สมร")
//การเข้าถึงแบบกำหนดช่วง
slice[low:high]

-Map
country := map[string]string{}
country["TH"] = "Thailand";
country["EN"] = "English";
----------------------------
population := map[string]int{}
population["Thailand"] = 70;
population["English"] = 55; 
----------------------------
coins :=map[string]string{"ETH":"Ether","BTC":"Bitcoin"}

-For  loop
โครงสร้าง  for  ค่าเริ่มต้นของตัวแปร; เงื่อนไข; เปลี่ยนแปลงค่าตัวแปร{
                คำสั่งเมื่อเงื่อนไขเป็นจริง
}

-function
โครงสร้าง มี3แบบ
    #แบบไม่มีรับและส่งค่า
    fucn ชื่อฟังก์ชั่น()){
        //คำสั่งต่าง
    }
    #แบบมีรับค่า
    fucn ชื่อฟังก์ชั่น(parameter1,parameter2,...){
        //คำสั่งต่าง
    }
    #แบบส่งค่าออกมา
    fucn ชื่อฟังก์ชั่น(){
        return ค่าที่ส่งออกไป
    }
    #แบบมีรับค่าและส่งออก
    fucn ชื่อฟังก์ชั่น(parameter1,parameter2,...){
        return ค่าที่ส่งออกไป
    }

-structure   การเอาข้อมูลต่างๆมารวมเข้าไว้ด้วยกัน
โครงสร้าง
type ชื่อสตรัคเจอร์ stuct{
    ตัวแปรที่1  ชนิดข้อมูลที่1;
    ตัวแปรที่2  ชนิดข้อมูลที่2;
}

