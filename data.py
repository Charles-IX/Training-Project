import mysql.connector
from mysql.connector import errorcode

try:
    # 数据库连接
    conn = mysql.connector.connect(
        host='localhost',  # 替换为你的主机名或IP地址
        user='charles',  # 替换为你的用户名
        password='1145141919',  # 替换为你的密码
        database='DrivTest'  # 替换为你的数据库名
    )

    cursor = conn.cursor()

    # 确保使用正确的数据库
    cursor.execute("USE DrivTest")

    # 创建 Questions 表，如果表不存在
    # cursor.execute('''
    # CREATE TABLE IF NOT EXISTS Questions (
    #     ID INT PRIMARY KEY,
    #     Question TEXT,
    #     Type TEXT,
    #     Chapter TEXT,
    #     `Key` TEXT,
    #     Answer TEXT,
    #     Times TEXT
    # )
    # ''')

    # 生成并插入数据
    for i in range(21, 151):
        X = chr(65 + (i % 4))  # 65是'A'的ASCII码
        question = f"测试用例-这道题的答案是{X}"
        key = str(ord(X) - 64)  # 1对应A, 2对应B, 3对应C, 4对应D
        chapter = str(2 + ((i - 1) // 20))
        
        cursor.execute('''
        INSERT INTO Questions (ID, Question, Type, Chapter, `Key`, Answer, Times)
        VALUES (%s, %s, %s, %s, %s, %s, %s)
        ''', (i, question, '1', chapter, key, question, '0'))

    # 提交事务并关闭连接
    conn.commit()
    conn.close()
    
except mysql.connector.Error as err:
    if err.errno == errorcode.ER_ACCESS_DENIED_ERROR:
        print("用户名或密码错误")
    elif err.errno == errorcode.ER_BAD_DB_ERROR:
        print("数据库不存在")
    else:
        print(err)
else:
    conn.close()
