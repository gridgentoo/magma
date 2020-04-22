Facebook объявил об открытии исходных текстов платформы Magma, включающей компоненты для быстрого развёртывания сотовых сетей (2G, 3G, 4G и 5G). Проект создан в рамках инициативы по обеспечению глобальной сетевой доступности и рассчитан на использование операторами связи. Платформа позволяет добиться кардинального упрощения процесса создания новых узлов сети, делая развёртывание LTE-сетей не сложнее, чем создание беспроводных точек доступа Wi-Fi. Код написан на языке Си и Python, и распространяется под лицензией BSD.

Оригинальный репозитарий: https://github.com/facebookincubator/magma  

Платформа также предлагает новый подход в работе операторов связи, основанный на использовании открытого ПО и позволяющий создавать новые типы сетей, в которых применяется быстрый цикл обновления и непрерывная интеграция программных компонентов. Централизованные компоненты для управления сетью можно размещать в приватных или общедоступных облачных окружениях. При этом платформа совместима с существующими базовыми станциями LTE и может сосуществовать и взаимодействовать с традиционными компонентами опорной сети LTE (core network).  

Платформа Magma уже прошла тестовое внедрение в компаниях Telefonica и BRCK, в которых применялась для охвата труднодоступных мест в странах Латинской Америки и для запуска новой LTE-сети в Кении. Платформа не предназначена для замены существующих внедрений EPC (Evolved Packet Core) и позиционируется для расширения существующих сервисов. Например, Magma позволяет ускорить развёртывание сотовой сети в сельской местности, а также может быть использована для создания частных сетей LTE или корпоративных беспроводных систем.  

Magma включает инструменты для автоматизации развёртывания сети, управляющее ПО и компоненты опорной сети для организации доставки пакетов. Для снижения сложности управления мобильными сетями в Magma предлагаются средства для автоматизации настройки, обновления ПО и добавления новых устройств. Открытый характер проекта позволяет операторам связи создавать решения не привязанные к одному поставщику оборудования, обеспечивает большую гибкость и предсказуемость, а также предоставляет больше возможностей для добавления новых сервисов и приложений. Для операторов, которые ограничены лицензированным диапазоном частот, Magma позволяет увеличить охват с помощью Wi-Fi и CBRS.

<b>AGW (Access Gateway)</b> - шлюз доступа, предоставляющий реализации PGW (Packet Data Network Gateway), SGW (Serving Gateway), MME (Mobility Management Entity) и AAA (Authentication, Authorization and Accounting). SGW обрабатывает и маршрутизирует пакеты для базовых станций. PGW обеспечивает подключение абонента к внешним сетям, выполняет фильтрацию пакетов и биллинг. MME обеспечивает мобильность, отслеживает перемещение абонента и выполняет миграцию между базовыми станциями. AAA предоставляет сетевые сервисы для аутентификации, авторизации и аккунтинга абонентов. Поддерживается работа с существующим оборудованием для сотовых сетей;  

Federation Gateway - шлюз для интеграции с опорной сетью мобильных операторов связи, использующий стандартные интерфейсы 3GPP для взаимодействия с существующими компонентами сети. Выполняет роль прокси между шлюзом доступа (AGW) и сетью оператора связи, обеспечивая работу таких функций, как аутентификация, списание средств, учёт и применение ограничений тарифных планов;  

Orchestrator - управляющий облачный сервис для настройки и мониторинга за беспроводной сетью, в том числе для анализа производительности сети и отслеживания потоков трафика. Для управления предлагается web-интерфейс. Orchestrator может запускаться в типовых облачных окружениях. Для взаимодействия с AGW и Federation Gateway применяется протокол gRPC, работающий поверх HTTP/2.  

# Magma

[![facebookincubator](https://circleci.com/gh/facebookincubator/magma.svg?style=shield)](https://circleci.com/gh/facebookincubator/magma)

Magma is an open-source software platform that gives network operators an open, flexible and extendable mobile core network solution. Magma enables better connectivity by:

* Allowing operators to offer cellular service without vendor lock-in with a modern, open source core network
* Enabling operators to manage their networks more efficiently with more automation, less downtime, better predictability, and more agility to add new services and applications
* Enabling federation between existing MNOs and new infrastructure providers for expanding rural infrastructure
* Allowing operators who are constrained with licensed spectrum to add capacity and reach by using Wi-Fi and CBRS


## Magma Architecture

The figure below shows the high-level Magma architecture. Magma is 3GPP generation (2G, 3G, 4G or upcoming 5G networks) and access network agnostic (cellular or WiFi). It can flexibly support a radio access network with minimal development and deployment effort.

Magma has three major components:

* **Access Gateway:** The Access Gateway (AGW) provides network services and policy enforcement. In an LTE network, the AGW implements an evolved packet core (EPC), and a combination of an AAA and a PGW. It works with existing, unmodified commercial radio hardware.

* **Orchestrator:** Orchestrator is a cloud service that provides a simple and consistent way to configure and monitor the wireless network securely. The Orchestrator can be hosted on a public/private cloud. The metrics acquired through the platform allows you to see the analytics and traffic flows of the wireless users through the Magma web UI.

* **Federation Gateway:** The Federation Gateway integrates the MNO core network with Magma by using standard 3GPP interfaces to existing MNO components.  It acts as a proxy between the Magma AGW and the operator's network and facilitates core functions, such as authentication, data plans, policy enforcement, and charging to stay uniform between an existing MNO network and the expanded network with Magma.

![Magma architecture diagram](docs/readmes/assets/magma_overview.png?raw=true "Magma Architecture")

## Usage Docs
The documentation for developing and using Magma is available at: [https://facebookincubator.github.io/magma](https://facebookincubator.github.io/magma)

## Join the Magma Community

- Mailing lists:
  - Join [magma-dev](https://groups.google.com/forum/#!forum/magma-dev) for technical discussions
  - Join [magma-announce](https://groups.google.com/forum/#!forum/magma-announce) for announcements
- Discord:
  - [magma\_dev](https://discord.gg/WDBpebF) channel

See the [CONTRIBUTING](CONTRIBUTING.md) file for how to help out.

## License

Magma is BSD License licensed, as found in the LICENSE file.
The EPC is OAI is offered under the OAI Apache 2.0 license, as found in the LICENSE file in the OAI directory.
