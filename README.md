<img align="right" src="https://layer5.io/assets/images/cube-sh-small.png" />

[![Docker Pulls](https://img.shields.io/docker/pulls/layer5/meshery.svg)](https://hub.docker.com/r/layer5/meshery)
[![Go Report Card](https://goreportcard.com/badge/github.com/layer5io/meshery)](https://goreportcard.com/report/github.com/layer5io/meshery)

# [Meshery](https://layer5.io/meshery)

A multi-mesh management plane with initial functionality as a playground to faciliate learning about functionality and performance of different service meshes. [Meshery](https://layer5.io/meshery) also serves as a service mesh performance benchmark.

- [Functionality](#functionality)
- [Running Meshery](#running)
  - [Quick start](https://layer5.io/meshery/#getting-started)
  - [Docs](https://meshery.layer5.io/docs)
- [Architecture](https://docs.google.com/presentation/d/1UbuYMpn-e-mWVYwEASy4dzyZlrSgZX6MUfNtokraT9o/edit?usp=sharing)
  - [Design document](https://docs.google.com/document/d/1nV8TunLmVC8j5cBELT42YfEXYmhG3ZqFtHxeG3-w9t0/edit?usp=sharing)
- [Contributing](CONTRIBUTING.md/#contributing)
  - [Write an adapter](CONTRIBUTING.md/#adapter)
  - [Build the project](CONTRIBUTING.md/#building)
  
In an effort to produce service mesh agnostic tooling, Meshery uses a [common performance benchmark specification](https://github.com/layer5io/service-mesh-benchmark-spec) to capture and share environment information and test configuration. 

## <a name="functionality">Functionality</a>
<img align="right" src="./ui/static/img/meshery.png?raw=true" alt="Service Mesh Playground" width="50%" />

### Multi-mesh Performance Benchmark

Meshery is intended to be a vendor and project-neutral utility for uniformly benchmarking the performance of service meshes. Between service mesh and proxy projects (and surprisingly, [within a single project](https://layer5.io/landscape#tools)), a number of different tools *and results* exist. 

### Multi-mesh Functionality Playground

A service mesh playground to faciliate learning about functionality of different service meshes. Meshery incorporates a visual interface for manipulating traffic routing rules. Sample applications will be included in Meshery. 

## <a name="running">Running Meshery</a>

See the [project site](https://layer5.io/meshery) for quick start instructions and [project documentation](https://meshery.layer5.io/docs) for a complete set of Meshery documentation.

## Presentations
- [KubeCon EU 2019](https://kccnceu19.sched.com/event/MPf7/service-meshes-at-what-cost-lee-calcote-layer5-girish-ranganathan-solarwinds?iframe=no&w=100%&sidebar=yes&bg=no) | ([video](https://www.youtube.com/watch?v=LxP-yHrKL4M&list=PLYjO73_1efChX9NuRaU7WocTbgrfvCoPE), [deck](https://calcotestudios.com/talks/decks/slides-kubecon-eu-2019-service-meshes-at-what-cost.html))
- Istio Founders Meetup @ KubeCon EU 2019 | [deck](https://calcotestudios.com/talks/decks/slides-istio-meetup-kubecon-eu-2019-istio-at-scale-large-and-small.html)
- [Cloud Native Rejekts EU 2019](https://cfp.cloud-native.rejekts.io/cloud-native-rejekts-eu-2019/speaker/GZQTEM/) | [deck](https://calcotestudios.com/talks/decks/slides-cloud-native-rejekts-2019-evaluating-service-meshes.html)
- [DockerCon 2019 Open Source Summit](https://dockercon19.smarteventscloud.com/connect/sessionDetail.ww?SESSION_ID=309149&tclass=popup#.XJxH-TOcbjI.twitter) | [deck](https://calcotestudios.com/talks/decks/slides-dockercon-2019-establishing-an-open-source-office.html), [video](https://www.docker.com/dockercon/2019-videos?watch=open-source-summit-service-mesh)
- [Container World 2019](https://tmt.knect365.com/container-world/speakers/lee-calcote) | [deck](https://calcotestudios.com/talks/decks/slides-container-world-2019-service-meshes-but-at-what-cost.html)
- [Service Mesh Day](https://servicemeshday.com/schedule/) | [deck](https://docs.google.com/presentation/d/1T0w5sXiUYtjHhmwJYF7VI-q5lgYAN46-Yn8ey0EZV-A/edit?usp=sharing), [video](https://youtu.be/CFj1O_uyhhs)
- [Innotech San Antonio](https://innotechsanantonio2019.sched.com/event/Lmlb/the-enterprise-path-to-service-mesh-architectures?iframe=no&w=100%&sidebar=yes&bg=no) | [deck](https://calcotestudios.com/talks/decks/slides-innotech-san-antonio-2019-the-enterprise-path-to-service-mesh.html)
- [CNCF Networking WG](https://github.com/cncf/wg-networking) | [deck](https://www.slideshare.net/leecalcote/benchmarking-service-meshes-cncf-networking-wg-141938576), [video](https://www.youtube.com/watch?v=2_JwCc-kLMA&list=PLYjO73_1efChX9NuRaU7WocTbgrfvCoPE)

<a name="contributing"></a><a name="community"></a>
## Community
<div>
This project is community-built and welcomes collaboration! 
<p>
  <ul>
    <li style="list-style-type: circle;"><em><strong>Join</strong></em> <a href="https://docs.google.com/document/d/1i6aa5OWTS7Ul1V2vJKVnJswlbBJi_CN5LhWs4xd85vw/edit#">weekly community meeting</a> on <a href="/assets/projects/meshery/Meshery-Community-Meeting.ics">Fridays from 2pm to 3pm Central</a>.</li>
    <ul>
        <li><em><strong>Watch</strong></em> community <a href="https://www.youtube.com/playlist?list=PL3A-A6hPO2IMPPqVjuzgqNU5xwnFFn3n0">meeting recordings</a>.</li>
    </ul>
    <li style="list-style-type: circle;"><em><strong>Access</strong></em> the <a href="https://drive.google.com/drive/folders/1cSSCn4428TFhyg5mCMZq-oDi-Gwrcqqt">community folder</a>.</li>
      <ul>
        <li><em><strong>Comment</strong></em> on the <a href="https://docs.google.com/document/d/1nV8TunLmVC8j5cBELT42YfEXYmhG3ZqFtHxeG3-w9t0/edit?usp=sharing">design document.</a></li>
      </ul>
    </ul>
</p>
</div>

## License

This repository and site are available as open source under the terms of the [Apache 2.0 License](https://opensource.org/licenses/Apache-2.0).

**About Layer5**
[Layer5.io](https://layer5.io) is a service mesh community, serving as a repository for information pertaining to the surrounding technology ecosystem (service meshes, api gateways, edge proxies, ingress and egress controllers) of microservice management in cloud native environments.
