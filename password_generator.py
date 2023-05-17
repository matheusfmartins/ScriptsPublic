def get_sufix():

    numbers = ["1","12","123","1234","12345","123456","12345678","123456789","1234567890"]
    dates = ["2000","2001","2002","2003","2004","2005","2006","2007","2008","2009","2010","2011","2012","2013","2014","2015","2016","2017","2018","2019","2020","2021","2022","2023"]

    specialChars = ["@", "!", "#", "$"]

    for i in range(len(numbers)):
        for j in range(len(specialChars)):
            numbers.append(specialChars[j] + numbers[i])

    for i in range(len(dates)):
        for j in range(len(specialChars)):
            dates.append(specialChars[j] + dates[i])

    final = numbers + dates

    return final

def get_prefix():
    return ["!","@","#","$"]

def recupera_palavras_chave():
    palavra = "Palavra para formar a senha"
    palavrasChave = []

    while palavra != "":
        palavra = input("Informe uma palavra chave (deixe em branco se não quiser mais inserir): ")

        if palavra == "":
            break

        palavrasChave.append(palavra)

    for i in range(len(palavrasChave)):
        palavra = palavrasChave[i]

        # A
        palavra = palavra.replace("a", "4")
        palavra = palavra.replace("A", "4")

        # E
        palavra = palavra.replace("e", "3")
        palavra = palavra.replace("E", "3")

        # I
        palavra = palavra.replace("i", "1")
        palavra = palavra.replace("I", "1")

        # O
        palavra = palavra.replace("o", "0")
        palavra = palavra.replace("O", "0")

        palavrasChave.append(palavra)

    return palavrasChave


def gerar_senhas(palavrasChave):
    senhas = palavrasChave

    sufix = get_sufix()
    prefix = get_prefix()

    # adiciona os prefixos nas palavras chave
    for i in range(len(senhas)):
        for j in range(len(prefix)):
            senhas.append(prefix[j] + senhas[i])

    # adiciona os sufixos nas palavras chave já com prefixo
    for i in range(len(senhas)):
        for j in range(len(sufix)):
            senhas.append(senhas[i] + sufix[j])

    return senhas

def gerar_palavras_admin():

    resp = input("Gerar senhas de admin? Y/N (Y):")

    if resp.upper() != "N":
        return ["admin", "admin1", "admin12", "admin123", "admin1234", "admin12345", "admin123456", 
        "!admin", "!admin1", "!admin12", "!admin123", "!admin1234", "!admin12345", "!admin123456",
        "admin!", "admin1!", "admin12!", "admin123!", "admin1234!", "admin12345!", "admin123456!",
        "admin@", "admin1@", "admin12@", "admin123@", "admin1234@", "admin12345@", "admin123456@",
        "admin@1", "admin@12", "admin@123", "admin@1234", "admin@12345", "admin@123456"]

    return []

def gerar_palavras_padrao():

    resp = input("Gerar senhas padrões? Y/N (Y):")

    if resp.upper() != "N":
        return ["senha", "123senha", "senha123", "senha1", "1senha", "senha12", "12senha", 
        "senha@123", "123mudar", "mudar123", "mudar@123", "Mudar123", "Mudar@123", "@Mudar123",
        "@mudar123", "mudar#123", "Mudar#123", "#Mudar123", "#mudar123", "mudar!123", "Mudar!123",
        "!Mudar123", "!mudar123", "minhasenha", "senhasecreta", "secreta", "secreto", "senhasupersecreta",
        "amesma", "a.mesma", "a@mesma", "a#mesma", "password", "123456", "1",
        "12", "123", "1234", "12345", "1234567", "12345678", "123456789",
        "1234567890", "21", "321", "4321", "54321", "654321", "7654321",
        "87654321", "987654321", "0987654321", "0p9o8i7u", "0p9o8i7u6y", "6y7u8i9o0p", "7u8i9o0p"
        "1q2w3e4r", "1q2w3e4r5t6y", "r4e3w2q1", "t5re3w2q1", "qwert", "qwerty", "qwertyu",
        "qwertyui", "qwertyuio", "qwertyuiop"]

    return []

def gerar_wordlist(senhas):
    f = open("custom_wordlist.txt", "w")

    for i in range(len(senhas)):

        if (i != 0):
            f.write("\n" + senhas[i])
        else:
            f.write(senhas[i])

    f.close()


print("--------------------------------------------------------------------------------")
print("Esse gerador de senhas irá gerar caracteres padrões com base nas palavras chaves")
print("--------------------------------------------------------------------------------")

palavrasChave = recupera_palavras_chave()
senhas = gerar_senhas(palavrasChave)
senhaAdmin = gerar_palavras_admin()
senhaPadrao = gerar_palavras_padrao()

senhas = senhaAdmin + senhaPadrao + senhas

#gera wordlist
gerar_wordlist(senhas)

print("[+] Wordlist gerada com sucesso: custom_wordlist.txt")
