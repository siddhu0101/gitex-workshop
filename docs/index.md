# Gitex Asia 2025: GitOps with ArgoCD

Welcome to our hands-on workshop at Gitex Asia 2025 in Singapore! In this session, hosted by Saiyam and Mumshad, you'll build a GitOps pipeline using ArgoCD on a KodeKloud Kubernetes playground. We'll deploy a Go-based web application, automate builds with GitHub Actions, and use ArgoCD to manage deployments in a declarative, Git-driven way.

## What is GitOps?

GitOps is a set of practices that use Git as the single source of truth for declarative infrastructure and application deployments. With tools like ArgoCD, changes to Git trigger automated deployments, ensuring consistency and enabling self-healing.

## Why ArgoCD?

ArgoCD is a Kubernetes-native GitOps tool that:
- Syncs applications from Git to clusters.
- Detects and corrects configuration drift.
- Provides a user-friendly UI for monitoring deployments.

## Prerequisites

- **KodeKloud Account**: Provided during the workshop for access to a pre-provisioned Kubernetes cluster.
- **GitHub Account**: To fork the workshop repo.
- **Basic Kubernetes Knowledge**: Familiarity with pods, services, and ingresses.

## Workshop Flow

### 1. Access the KodeKloud Playground

1. Log into the KodeKloud playground (credentials provided).
2. Verify cluster access:
   ```bash
   kubectl get nodes