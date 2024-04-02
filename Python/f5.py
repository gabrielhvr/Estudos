import time
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC

# Configuração do navegador
driver = webdriver.Chrome()  # Certifique-se de ter o chromedriver instalado e no PATH
url = "https://bloodsailsociety.wordpress.com/"

# Função para atualizar a página
def refresh_page():
    driver.refresh()
    print("Página atualizada.")

# Função para verificar se houve mudança na área específica do site
def check_for_change_in_specific_area():
    # Espera até que o elemento esteja presente na página
    try:
        area_element = WebDriverWait(driver, 10).until(EC.presence_of_element_located((By.CSS_SELECTOR, "body")))
        area_content = area_element.text
        
        # Compare o conteúdo da área com o conteúdo anterior
        if area_content != previous_area_content:
            print("Houve uma mudança na área específica do site!")
            # Se houver uma mudança, você pode fazer o que for necessário aqui
            # Por exemplo, enviar um email, registrar a mudança em um arquivo de log, etc.
        
        # Atualize o conteúdo anterior com o conteúdo atual para a próxima verificação
        return area_content
    except Exception as e:
        print("Erro ao encontrar o elemento:", e)
        return None

# Tempo entre cada verificação (em segundos)
intervalo = 60  # 1 minuto

try:
    driver.get(url)
    previous_area_content = check_for_change_in_specific_area()
    while True:
        refresh_page()
        time.sleep(intervalo)
        previous_area_content = check_for_change_in_specific_area()
except KeyboardInterrupt:
    print("Script interrompido pelo usuário.")
finally:
    driver.quit()


